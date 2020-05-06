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

	r.HandleFunc("/rest/api/v1/cms/home", controllers.GetCmsHomeData).Methods("GET")

	r.HandleFunc("/rest/api/v1/jobs/index", controllers.GetJobs).Queries("timestamp", "{[0-9]+}").Methods("GET")
	r.HandleFunc("/rest/api/v1/jobs/search/location", controllers.GetJobsByLocation).Queries("cursor", "{cursor:[0-9]*$}", "location", "{location:[a-zA-Z ]*$}").Methods("GET")
	r.HandleFunc("/rest/api/v1/jobs/search/experience", controllers.GetJobsByExperience).Queries("cursor", "{cursor:[0-9]*$}", "experience", "{experience:[a-zA-Z ]*$}").Methods("GET")
	// r.HandleFunc("/rest/api/v1/jobs/index", controllers.GetJobs).Queries("cursor", "{[0-9]+}").Methods("GET")

	r.HandleFunc("/rest/api/v1/jobs/", controllers.CreateJob).Methods("POST")

	r.HandleFunc("/rest/api/v1/jobs/company/search/{companyUUID:[A-Za-z0-9_@./#&+-]*$}", controllers.GetJobsByCompany).Methods("GET")
	r.HandleFunc("/rest/api/v1/jobs/company/list/", controllers.GetCompanyList).Methods("GET")

	r.HandleFunc("/rest/api/v1/jobs/company/location/{location:[A-Za-z ]*$}", controllers.GetCompaniesByLocation).Methods("GET")
	r.HandleFunc("/rest/api/v1/jobs/company/company/{company:[A-Za-z0-9 ]*$}", controllers.GetLocationsByCompany).Methods("GET")

	c.SubscribeTopics([]string{"new_job", "del_job", "job_location"}, nil)
	go http.ListenAndServe(":8080", r)

	for {
		msg, err := c.ReadMessage(-1)

		if err == nil {
			s := strings.Split(msg.TopicPartition.String(), "[")
			topic, _ := s[0], s[1]
			fmt.Println(topic)
			switch topic {
			case "new_job":
				fmt.Println("new job")
				consumergroups.Addnewjob(msg)
				break
			case "job_location":
				fmt.Println("new job location")
				consumergroups.AddNewJobLocation(msg)
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
