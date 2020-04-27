package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/carlosbenavides123/DevJobs/MainServer/mvc/services"
)

func GetCompaniesByLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	res, apiErr := services.GetCompaniesByLocation(r)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write([]byte(jsonValue))
		return
	}
	jsonValue, _ := json.Marshal(res)
	w.Write(jsonValue)
}
