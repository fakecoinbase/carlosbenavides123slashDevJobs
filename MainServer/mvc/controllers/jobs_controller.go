package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/carlosbenavides123/DevJobs/MainServer/mvc/services"
)

func GetJobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Content-Type", "application/json")
	jobs, apiErr := services.GetJobs()
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write([]byte(jsonValue))
		return
	}
	jsonValue, _ := json.Marshal(jobs)
	w.Write(jsonValue)
}
