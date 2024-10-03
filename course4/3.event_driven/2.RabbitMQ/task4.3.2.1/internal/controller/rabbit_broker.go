package controller

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

var m = make(map[string]*amqp.Channel)

func GetRabbitBroker() *amqp.Channel {
	if val, ok := m["rabbit"]; ok {
		return val
	}
	var conn *amqp.Connection
	var err error
	i := 0
	maxAttemps := 10
	for conn, err = amqp.Dial("amqp://rabbitmq:rabbitmq@rabbitmq:5672/"); err != nil; conn, err = amqp.Dial("amqp://rabbitmq:rabbitmq@rabbitmq:5672/") {
		log.Printf("try:%d:%v", i, err)
		i++
		time.Sleep(3 * time.Second)
		if i > maxAttemps {
			log.Fatal("rabbit broker connect timeout", err)
		}
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Println("Failed to open a channel", err)
	}
	_, err = ch.QueueDeclare(
		"my_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("Failed to declare a queue", err)
	}
	m["rabbit"] = ch
	return ch
}
