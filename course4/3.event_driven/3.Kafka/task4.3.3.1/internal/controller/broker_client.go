package controller

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/streadway/amqp"
)

func GetBroker(broker string) Broker {
	switch broker {
	case "kafka":
		return NewKafkaAdapter(GetKafkaWriter())
	case "rabbit":
		return NewRabbitAdapter(GetRabbitBroker())
	default:
		log.Fatal("broker not supported")
	}
	return nil
}

func GetRabbitBroker() *amqp.Channel {
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
	log.Println("rabbit start")
	return ch
}

type RabbitAdapter struct {
	ch *amqp.Channel
}

func NewRabbitAdapter(ch *amqp.Channel) *RabbitAdapter {
	return &RabbitAdapter{ch: ch}
}

func (ra *RabbitAdapter) Publish(msg []byte) error {
	return ra.ch.Publish(
		"",
		"my_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)
}

func GetKafkaWriter() *kafka.Writer {
	topic := "my_topic"

	writer := &kafka.Writer{
		Addr:     kafka.TCP("kafka1:9092"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	log.Println("kafka start writer")
	return writer
}

type KafkaAdapter struct {
	writer *kafka.Writer
}

func NewKafkaAdapter(writer *kafka.Writer) *KafkaAdapter {
	return &KafkaAdapter{
		writer: writer,
	}
}

func (ka *KafkaAdapter) Publish(msg []byte) error {
	return ka.writer.WriteMessages(context.Background(), kafka.Message{
		Value: msg,
	})
}
