package kafkaconf

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func NewProducer() (p *kafka.Producer) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "64.255.121.38:19092"})
	if err != nil {
		panic(err.Error())
	}
	return p
}

func NewConsumer() (c *kafka.Consumer) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "64.255.121.38:19092",
		"group.id":          "CMS",
	})
	if err != nil {
		panic(err)
	}
	return c
}
