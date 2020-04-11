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

	res, err := db.Query(`select j.job_uuid, j.job_title, j.job_link, j.job_posted, j.job_found, j.job_idx, c.name, c.cloudinary, l.job_level
	from companies c
	inner join jobs_pivot jp on jp.company_uuid = c.company_uuid
	inner join jobs j on jp.job_uuid = j.job_uuid
	inner join levels l on j.experience_level = l.id
	ORDER BY j.job_posted DESC`)
	if err != nil {
		panic(err.Error())
	}
	job := []*Job{}

	for res.Next() {
		var JobUUID, JobTitle, JobLink string
		var JobPosted, JobFound, JobIdx int64
		var CompanyName, Cloudinary, JobLevel string

		err = res.Scan(&JobUUID, &JobTitle, &JobLink, &JobPosted, &JobFound, &JobIdx, &CompanyName, &Cloudinary, &JobLevel)
		if err != nil {
			panic(err.Error())
		}
		jobRow := &Job{}
		jobRow.JobUUID = JobUUID
		jobRow.JobTitle = JobTitle
		jobRow.JobLink = JobLink
		jobRow.JobPosted = JobPosted
		jobRow.JobFound = JobFound
		jobRow.JobIdx = JobIdx
		jobRow.CompanyName = CompanyName
		jobRow.Cloudinary = Cloudinary
		jobRow.JobLevel = JobLevel

		job = append(job, jobRow)
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
