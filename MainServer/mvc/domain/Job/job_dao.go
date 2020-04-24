package job

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/carlosbenavides123/DevJobs/MainServer/dbconf"
	"github.com/carlosbenavides123/DevJobs/MainServer/mvc/utils"
	uuid "github.com/nu7hatch/gouuid"
)

func GetJobs() ([]*Job, *utils.ApplicationError) {

	db := dbconf.DbConn()
	defer db.Close()

	res, err := db.Query(`select j.job_uuid, j.job_title, j.job_link, j.job_location, j.job_posted, j.job_found, j.job_idx, c.company_name, c.company_cloudinary, l.job_level
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
		var JobUUID, JobTitle, JobLink, JobLocation string
		var JobPosted, JobFound, JobIdx int64
		var CompanyName, Cloudinary, JobLevel string

		err = res.Scan(&JobUUID, &JobTitle, &JobLink, &JobLocation, &JobPosted, &JobFound, &JobIdx, &CompanyName, &Cloudinary, &JobLevel)
		if err != nil {
			panic(err.Error())
		}
		jobRow := &Job{}
		jobRow.JobUUID = JobUUID
		jobRow.JobTitle = JobTitle
		jobRow.JobLink = JobLink
		jobRow.JobLocation = JobLocation
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

func CreateJob(newCompany *NewCompany) (*NewCompany, *utils.ApplicationError) {
	db := dbconf.DbConn()
	defer db.Close()

	res, dbPrepareErr := db.Prepare(`INSERT INTO companies(company_uuid, company_name, company_cloudinary) 
				VALUES (?, ?, ?)
				`)
	uuid, dbPrepareErr := uuid.NewV4()
	if dbPrepareErr != nil {
		panic(error(dbPrepareErr))
	}
	res.Exec(uuid.String(), newCompany.CompanyName, newCompany.Cloudinary)

	db2 := dbconf.DbConnToScrappy()
	defer db2.Close()
	res2, dbPrepareErr2 := db2.Prepare(`INSERT INTO companies(company_uuid, company_name, company_scrape_website, greenhouse, lever, other)
										VALUES(?, ?, ?, ?, ?, ?)`)
	if dbPrepareErr2 != nil {
		panic(error(dbPrepareErr2))
	}
	var greenhouse, lever, other bool
	if strings.Contains(newCompany.CompanyWebsite, "greenhouse") {
		greenhouse = true
	} else if strings.Contains(newCompany.CompanyWebsite, "lever") {
		lever = true
	} else {
		other = true
	}
	res2.Exec(uuid.String(), newCompany.CompanyName, newCompany.CompanyWebsite, greenhouse, lever, other)
	return newCompany, nil
}

func GetJobsByCompany(companyUUID string) ([]*Job, *utils.ApplicationError) {
	db := dbconf.DbConn()
	defer db.Close()

	stmt, dbPrepareErr := db.Prepare(`select j.job_uuid, j.job_title, j.job_link, j.job_location, j.job_posted, j.job_found, j.job_idx, c.company_name, c.company_cloudinary, l.job_level
		from companies c inner join jobs_pivot jp on jp.company_uuid = c.company_uuid 
		AND c.company_uuid=?
		inner join jobs j on jp.job_uuid = j.job_uuid 
		inner join levels l on j.experience_level = l.id 
		ORDER BY j.job_posted DESC`)

	if dbPrepareErr != nil {
		panic(error(dbPrepareErr))
	}

	res, queryErr := stmt.Query(companyUUID)

	if queryErr != nil {
		panic(error(queryErr))
	}

	job := []*Job{}

	for res.Next() {
		var JobUUID, JobTitle, JobLink, JobLocation string
		var JobPosted, JobFound, JobIdx int64
		var CompanyName, Cloudinary, JobLevel string

		scanErr := res.Scan(&JobUUID, &JobTitle, &JobLink, &JobLocation, &JobPosted, &JobFound, &JobIdx, &CompanyName, &Cloudinary, &JobLevel)

		if scanErr != nil {
			panic(scanErr.Error())
		}

		jobRow := &Job{}
		jobRow.JobUUID = JobUUID
		jobRow.JobTitle = JobTitle
		jobRow.JobLink = JobLink
		jobRow.JobLocation = JobLocation
		jobRow.JobPosted = JobPosted
		jobRow.JobFound = JobFound
		jobRow.JobIdx = JobIdx
		jobRow.CompanyName = CompanyName
		jobRow.Cloudinary = Cloudinary
		jobRow.JobLevel = JobLevel

		job = append(job, jobRow)
	}
	return job, nil
}

func GetCompanyList() ([]*Company, *utils.ApplicationError) {
	db := dbconf.DbConn()
	defer db.Close()

	res, queryErr := db.Query(`select c.company_name, c.company_uuid from companies c`)
	if queryErr != nil {
		panic(queryErr.Error())
	}
	companies := []*Company{}
	for res.Next() {
		var CompanyName, CompanyUUID string

		scanErr := res.Scan(&CompanyName, &CompanyUUID)

		if scanErr != nil {
			panic(scanErr.Error())
		}
		companyRow := &Company{}
		companyRow.CompanyName = CompanyName
		companyRow.CompanyUUID = CompanyUUID
		companies = append(companies, companyRow)
	}
	return companies, nil
}
