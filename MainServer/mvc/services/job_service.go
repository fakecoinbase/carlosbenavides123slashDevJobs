package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	job "github.com/carlosbenavides123/DevJobs/MainServer/mvc/domain/Job"
	"github.com/carlosbenavides123/DevJobs/MainServer/mvc/utils"
	"github.com/gorilla/mux"
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

func GetJobsByCompany(r *http.Request) ([]*job.Job, *utils.ApplicationError) {
	params := mux.Vars(r)
	companyUUID := params["companyUUID"]
	if companyUUID == "" {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Company is required!"),
			StatusCode: http.StatusNotFound,
			Code:       "No Jobs available.",
		}
	}
	return job.GetJobsByCompany(companyUUID)
}

func GetCompanyList() ([]*job.Company, *utils.ApplicationError) {
	return job.GetCompanyList()
}
