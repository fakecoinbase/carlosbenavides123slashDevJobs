package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/carlosbenavides123/DevJobs/MainServer/mvc/services"
	"github.com/golang/protobuf/jsonpb"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func GetCmsHomeData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	res := services.GetCmsHomeData(r)
	if res == nil {

	} else {
		m := jsonpb.Marshaler{}
		result, _ := m.MarshalToString(res)
		w.Write([]byte(result))
	}
}

func GetCmsCompanyData(p *kafka.Producer, c *kafka.Consumer, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	result := services.GetCmsCompanyData(p, c, r)
	jsonValue, _ := json.Marshal(result)
	w.Write([]byte(jsonValue))
}

func ReverseRequestWithKafka(p *kafka.Producer, c *kafka.Consumer, f func(p *kafka.Producer, c *kafka.Consumer, w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { f(p, c, w, r) })
}
