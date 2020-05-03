package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	job "github.com/carlosbenavides123/DevJobs/MainServer/mvc/domain/Job"
	"github.com/carlosbenavides123/DevJobs/MainServer/mvc/utils"
	"github.com/gorilla/mux"
)

func GetJobs(r *http.Request) (*job.JobResponse, *utils.ApplicationError) {
	jobIdx := r.URL.Query().Get("timestamp")
	// params := mux.Vars(r)
	// timestamp := params["timestamp"]
	if jobIdx == "" {
		jobIdx = "2147000"
	}

	fmt.Println(jobIdx)

	// if timestamp == "" {
	// 	return nil, &utils.ApplicationError{
	// 		Message:    fmt.Sprintf("timestamp parameter required!"),
	// 		StatusCode: http.StatusNotFound,
	// 		Code:       "No Jobs available.",
	// 	}
	// }
	return job.GetJobs(jobIdx)
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

func GetJobsByLocation(r *http.Request) (*job.JobResponse, *utils.ApplicationError) {
	jobIdx := r.URL.Query().Get("cursor")
	loc := r.URL.Query().Get("location")
	if loc == "" {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Location is required!"),
			StatusCode: http.StatusBadRequest,
			Code:       "Location Missing.",
		}
	}
	if jobIdx == "" {
		jobIdx = "2147000"
	}
	return job.GetJobsByLocation(loc, jobIdx)
}

func GetJobsByExperience(r *http.Request) (*job.JobResponse, *utils.ApplicationError) {
	cursor := r.URL.Query().Get("cursor")
	experience := r.URL.Query().Get("experience")
	if experience == "" {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Experience is required!"),
			StatusCode: http.StatusBadRequest,
			Code:       "Experience Missing.",
		}
	}
	if cursor == "" {
		cursor = "2147000"
	}
	return job.GetJobsByExperience(experience, cursor)
}
