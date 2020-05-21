package cms

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/carlosbenavides123/DevJobs/MainServer/kafkaconf"
	"github.com/carlosbenavides123/DevJobs/MainServer/mvc/utils"
	"github.com/carlosbenavides123/DevJobs/MainServer/pb/company/companycmspb"
	"github.com/carlosbenavides123/DevJobs/MainServer/pb/company/companypb"
	"github.com/carlosbenavides123/DevJobs/MainServer/pb/company/companyrequestpb"
	"github.com/carlosbenavides123/DevJobs/MainServer/pb/company/updatecompanypb"
	"github.com/golang/protobuf/proto"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func GetCmsHomeData() (companyResponse *companypb.CompanyResponse) {
	c := kafkaconf.NewConsumer()
	c.SubscribeTopics([]string{"ResponseCmsHome"}, nil)

	p := kafkaconf.NewProducer()
	defer p.Close()
	topic := "RequestCmsHome"

	produceKafkaMessage(p, topic, []byte("Lol"))

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Println("error sending message")
					produceKafkaMessage(p, topic, []byte("lol"))
				} else {
					fmt.Println("message was sent successfully")
				}
			}
		}
	}()

	run := true

	for run == true {
		ev := c.Poll(15)
		if ev == nil {
			continue
		}

		switch e := ev.(type) {
		case *kafka.Message:
			s := strings.Split(e.TopicPartition.String(), "[")
			topic, _ := s[0], s[1]
			fmt.Println(topic)
			if topic == "ResponseCmsHome" {
				run = false
				companies := &companypb.CompanyResponse{}
				if err := proto.Unmarshal(e.Value, companies); err != nil {
					log.Fatalln("Failed to parse Job:", err)
				}
				return companies
			}
			fmt.Printf("%% Message on %s:\n%s\n",
				e.TopicPartition, string(e.Value))
			if e.Headers != nil {
				fmt.Printf("%% Headers: %v\n", e.Headers)
			}
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
			if e.Code() == kafka.ErrAllBrokersDown {
				run = false
			}
		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}
	return nil
}

func GetCmsCompanyData(p *kafka.Producer, c *kafka.Consumer, company string) *CompanyCms {
	topicProduce := "RequestCMSCompany"
	topicConsume := "ResponseCompanyCMS"
	c.SubscribeTopics([]string{topicConsume}, nil)

	companyrequestpb := &companyrequestpb.CompanyRequest{
		CompanyName: company,
	}
	companyrequest, err := proto.Marshal(companyrequestpb)
	if err != nil {
		panic(err.Error())
	}

	go produceKafkaMessage(p, topicProduce, companyrequest)

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Println("error sending message")
					produceKafkaMessage(p, topicProduce, companyrequest)
				} else {
					fmt.Println("message was sent successfully")
				}
			}
		}
	}()

	run := true
	companyresponse := &companycmspb.CompanyCmsDetails{}
	for run == true {
		msg, err := c.ReadMessage(-1)

		if err == nil {
			s := strings.Split(msg.TopicPartition.String(), "[")
			topic, _ := s[0], s[1]
			fmt.Println(topic)
			switch topic {
			case "ResponseCompanyCMS":
				c.Commit()
				if err := proto.Unmarshal(msg.Value, companyresponse); err != nil {
					log.Fatalln("Failed to parse Job:", err)
				}
				if companyresponse.CompanyName == company {
					run = false
				}
				break
			default:
				break
			}
		}
	}

	companycms := &CompanyCms{}
	companycms.CompanyName = companyresponse.CompanyName
	companycms.CompanyWebsite = companyresponse.CompanyWebsite
	companycms.GreenHouse = companyresponse.GreenHouse
	companycms.Lever = companyresponse.Lever
	companycms.Other = companyresponse.Other
	companycms.WantedDepartments = companyresponse.WantedDepartments
	companycms.WantedLocations = companyresponse.WantedLocations

	return companycms
}

func UpdateCompanyCMSData(updateCompanyDetails *CompanyCms, p *kafka.Producer) (*utils.ApplicationSuccess, *utils.ApplicationError) {
	topicProduce := "RequestUpdateCms"

	companyCmsRequest := &updatecompanypb.UpdateCompanyDetails{
		CompanyUUID:       updateCompanyDetails.CompanyUUID,
		CompanyName:       updateCompanyDetails.CompanyName,
		CompanyWebsite:    updateCompanyDetails.CompanyWebsite,
		GreenHouse:        updateCompanyDetails.GreenHouse,
		Lever:             updateCompanyDetails.Lever,
		Other:             updateCompanyDetails.Other,
		WantedDepartments: updateCompanyDetails.WantedDepartments,
		WantedLocations:   updateCompanyDetails.WantedLocations,
	}

	companyrequest, err := proto.Marshal(companyCmsRequest)

	if err != nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Form has bad input!"),
			StatusCode: http.StatusBadRequest,
			Code:       "Malformed request",
		}
	}

	produceKafkaMessage(p, topicProduce, companyrequest)

	return &utils.ApplicationSuccess{
		Message:    fmt.Sprintf("Update has been received!"),
		StatusCode: http.StatusAccepted,
		Code:       "Updated",
	}, nil
}

func produceKafkaMessage(p *kafka.Producer, topic string, data []byte) {
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)
}
