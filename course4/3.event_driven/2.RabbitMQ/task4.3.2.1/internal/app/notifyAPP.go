package app

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

type Message struct {
	Text  string `json:"text"`
	Email string `json:"email"`
}

func RunNotifyApp() error {
	var (
		conn *amqp.Connection
		err  error
	)

	i := 0
	maxAttemps := 10

	for conn, err = amqp.Dial("amqp://rabbitmq:rabbitmq@rabbitmq:5672/"); err != nil; conn, err = amqp.Dial("amqp://rabbitmq:rabbitmq@rabbitmq:5672/") {
		log.Printf("try:%d:%v", i, err)
		i++
		time.Sleep(3 * time.Second)
		if i > maxAttemps {
			return fmt.Errorf("rabbit broker connect timeout %w", err)
		}
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"my_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	log.Println("notify service start")
	for {
		msgs, err := ch.Consume(
			queue.Name,
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Println(err)
		}
		for msg := range msgs {
			m := Message{}
			err := json.Unmarshal(msg.Body, &m)
			if err != nil {
				log.Println(err)
			} else {
				// отправка сообщения пользователю на почту
				log.Printf("send to %s, msg: %s", m.Email, m.Text)
			}
		}
	}
}
