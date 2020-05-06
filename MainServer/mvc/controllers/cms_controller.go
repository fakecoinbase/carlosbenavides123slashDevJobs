package controllers

import (
	"net/http"

	"github.com/carlosbenavides123/DevJobs/MainServer/mvc/services"
	"github.com/golang/protobuf/jsonpb"
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
