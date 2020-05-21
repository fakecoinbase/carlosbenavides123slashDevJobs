package kafkaconf

import (
	"os"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func NewProducer() (p *kafka.Producer) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": os.Getenv("KAFKA_SERVER_IP")})
	if err != nil {
		panic(err.Error())
	}
	return p
}

func NewConsumer() (c *kafka.Consumer) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_SERVER_IP"),
		"group.id":          "CMS",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	return c
}
