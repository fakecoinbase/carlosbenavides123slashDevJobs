package job

import (
	"fmt"
	"net/http"

	"github.com/carlosbenavides123/DevJobs/MainServer/dbconf"
	"github.com/carlosbenavides123/DevJobs/MainServer/mvc/utils"
)

func GetJobs() ([]*Job, *utils.ApplicationError) {

	db := dbconf.DbConn()
	defer db.Close()

	res, err := db.Query(`SELECT * FROM jobs`)
	if err != nil {
		panic(err.Error())
	}
	job := []*Job{}

	for res.Next() {
		var JobID, CompanyName, JobLink string
		var JobPosted, JobFound int64
		var Active bool
		var time []uint8

		err = res.Scan(&JobID, &CompanyName, &JobLink, &JobPosted, &JobFound, &Active, &time)
		if err != nil {
			panic(err.Error())
		}
		temp_job := &Job{}
		temp_job.UUID = JobID
		temp_job.CompanyName = CompanyName
		temp_job.JobLink = JobLink
		temp_job.JobPosted = JobPosted
		temp_job.JobFound = JobFound

		job = append(job, temp_job)
	}
	if err == nil {
		return job, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("Jobs were not found"),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}
}
