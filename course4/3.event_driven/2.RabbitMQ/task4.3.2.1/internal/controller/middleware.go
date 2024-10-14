package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/streadway/amqp"
	"go.uber.org/ratelimit"

	"metrics/internal/API/gRPCAuth"
)

type Message struct {
	Email string `json:"email"`
	Text  string `json:"text"`
}

type Broker interface {
	Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
}

type Middleware struct {
	auth         gRPCAuth.AuthServiceClient
	countRequest map[string]int
	broker       Broker
}

func NewMiddleware(auth gRPCAuth.AuthServiceClient, broker Broker) *Middleware {
	mid := &Middleware{
		auth:         auth,
		countRequest: make(map[string]int),
		broker:       broker,
	}
	go mid.CheckingCount()

	return mid
}

func (m *Middleware) CheckingCount() {
	maxCount := 5
	rl := ratelimit.New(5, ratelimit.Per(time.Minute))
	for {
		rl.Take()
		for key, val := range m.countRequest {
			if val >= maxCount {
				m.countRequest[key] = maxCount
			} else {
				m.countRequest[key] = val + 1
			}
		}
	}
}

func (m *Middleware) AuthVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := getToken(r)
		_, err := m.auth.VerifyToken(context.Background(), token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(err.Error()))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func (m *Middleware) Limiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := getToken(r)
		user, _ := m.auth.VerifyToken(context.Background(), token)
		if _, ok := m.countRequest[user.Email]; !ok {
			m.countRequest[user.Email] = 5
		}
		if m.countRequest[user.Email] <= 0 {
			w.WriteHeader(http.StatusTooManyRequests)
			// send to rabbit message
			message := Message{Email: user.Email, Text: "You can do only 5 request per minute, sorry"}
			msg, _ := json.Marshal(message)
			err := m.broker.Publish("", "my_queue", false, false, amqp.Publishing{
				ContentType: "text/plain",
				Body:        msg,
			})
			if err != nil {
				log.Println(err)
			}
		} else {
			m.countRequest[user.Email]--
			next.ServeHTTP(w, r)
		}
	})
}

func getToken(r *http.Request) *gRPCAuth.Token {
	tokenStr := r.Header.Get("Authorization")
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
	token := gRPCAuth.Token{Token: tokenStr}

	return &token
}
