package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/carlosbenavides123/test_kafka_consumer/jobpb"
	proto "github.com/golang/protobuf/proto"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "192.168.1.66:19092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"foo"}, nil)

	for {
		msg, err := c.ReadMessage(-1)

		if err == nil {
			s := strings.Split(msg.TopicPartition.String(), "[")
			topic, _ := s[0], s[1]

			switch topic {
			case "foo":
				fmt.Println("lmao")
			default:
				break
			}

			job := &jobpb.Job{}
			if err := proto.Unmarshal(msg.Value, job); err != nil {
				log.Fatalln("Failed to parse Job:", err)
			}

			fmt.Println(proto.MarshalTextString(job))

			// fmt.Printf(job)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
