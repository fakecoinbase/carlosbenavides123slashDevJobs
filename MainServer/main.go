package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/carlosbenavides123/DevJobs/MainServer/consumergroups"
	"github.com/carlosbenavides123/DevJobs/MainServer/mvc/controllers"
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

	r := mux.NewRouter()

	// websocket for future

	// rest apis
	r.HandleFunc("/rest/api/v1/jobs/", controllers.GetJobs).Methods("GET")
	r.HandleFunc("/rest/api/v1/jobs/", controllers.CreateJob).Methods("POST")

	c.SubscribeTopics([]string{"new_job", "del_job"}, nil)
	go http.ListenAndServe(":8080", r)

	for {
		msg, err := c.ReadMessage(-1)

		if err == nil {
			s := strings.Split(msg.TopicPartition.String(), "[")
			topic, _ := s[0], s[1]

			switch topic {
			case "new_job":
				fmt.Println("lmao")
				consumergroups.Addnewjob(msg)
				break
			default:
				break
			}

			// job := &jobpb.Job{}
			// if err := proto.Unmarshal(msg.Value, job); err != nil {
			// 	log.Fatalln("Failed to parse Job:", err)
			// }

			// fmt.Println(proto.MarshalTextString(job))

			// fmt.Printf(job)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
