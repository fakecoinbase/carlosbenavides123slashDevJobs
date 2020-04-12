package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	job "github.com/carlosbenavides123/DevJobs/MainServer/mvc/domain/Job"
	"github.com/carlosbenavides123/DevJobs/MainServer/mvc/utils"
)

func GetJobs() ([]*job.Job, *utils.ApplicationError) {
	return job.GetJobs()
}

func CreateJob(r *http.Request) (*job.NewCompany, *utils.ApplicationError) {
	decoder := json.NewDecoder(r.Body)
	var newCompany *job.NewCompany
	decodeErr := decoder.Decode(&newCompany)
	if decodeErr != nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Missing one or more paramters."),
			StatusCode: http.StatusBadRequest,
			Code:       "Malformed request",
		}
	}
	fmt.Printf("%+v\n", newCompany)
	return job.CreateJob(newCompany)
}
