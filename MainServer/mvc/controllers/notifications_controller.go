package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/carlosbenavides123/DevJobs/MainServer/mvc/services"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func GetNotificationPreference(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("ping")
	res, apiErr := services.GetNotificationPreference(r)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write([]byte(jsonValue))
		return
	}
	jsonValue, _ := json.Marshal(res)
	w.Write(jsonValue)
}

func CreateNotificationPreference(p *kafka.Producer, c *kafka.Consumer, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	services.CreateNotificationPreference(r, p)
	w.Write([]byte("hello"))
}

func UpdateNotificationPreference(p *kafka.Producer, c *kafka.Consumer, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("ping_update")
	res, apiErr := services.UpdateNotificationPreference(r, p)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write([]byte(jsonValue))
		return
	}
	jsonValue, _ := json.Marshal(res)
	w.Write(jsonValue)

}
