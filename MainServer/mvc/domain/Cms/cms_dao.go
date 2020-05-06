package cms

import (
	"fmt"
	"log"
	"strings"

	"github.com/carlosbenavides123/DevJobs/MainServer/kafkaconf"
	"github.com/carlosbenavides123/DevJobs/MainServer/pb/company/companypb"
	"github.com/golang/protobuf/proto"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func GetCmsHomeData() (companyResponse *companypb.CompanyResponse) {
	c := kafkaconf.NewConsumer()
	c.SubscribeTopics([]string{"ResponseCmsHome"}, nil)

	p := kafkaconf.NewProducer()
	defer p.Close()
	topic := "RequestCmsHome"

	produceKafkaMessage(p, topic)

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Println("error sending message")
					produceKafkaMessage(p, topic)
				} else {
					fmt.Println("message was sent successfully")
				}
			}
		}
	}()

	for true {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			s := strings.Split(msg.TopicPartition.String(), "[")
			topic, _ := s[0], s[1]
			fmt.Println(topic)
			if topic == "ResponseCmsHome" {
				companies := &companypb.CompanyResponse{}
				if err := proto.Unmarshal(msg.Value, companies); err != nil {
					log.Fatalln("Failed to parse Job:", err)
				}
				return companies
			}
		} else {
			break
		}
	}
	return nil
}

func produceKafkaMessage(p *kafka.Producer, topic string) {
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("ping"),
	}, nil)
}
