package broker_server

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"

	"metrics/internal/API/gRPCNotify"
)

func StartKafka(NotifyCl gRPCNotify.NotifyServiceClient) {
	topic := "my_topic"
	partition := 0
	config := kafka.ReaderConfig{
		Brokers:   []string{"kafka1:9092"},
		Topic:     topic,
		Partition: partition,
		MinBytes:  10e3,
		MaxBytes:  10e6,
		MaxWait:   1 * time.Second,
	}
	i := 0
	maxAttemps := 10
	reader := &kafka.Reader{}
	for reader = kafka.NewReader(config); reader == nil; reader = kafka.NewReader(config) {
		log.Printf("try %d", i)
		i++
		if i > maxAttemps {
			log.Fatal("can not connect to Kafka")
		}
		time.Sleep(3 * time.Second)
	}
	defer reader.Close()
	log.Println("starting kafka reader")
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
			reader = kafka.NewReader(config)
		}
		m := Message{}
		err = json.Unmarshal(msg.Value, &m)
		if err != nil {
			log.Println(err)
		} else {
			// отправка сообщения пользователю на почту
			_, err = NotifyCl.SendMessage(context.Background(), &gRPCNotify.Message{Email: m.Email, Text: m.Text})
			if err != nil {
				log.Println(err)
			}
		}
	}
}
