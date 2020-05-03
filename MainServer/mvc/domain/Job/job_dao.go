package job

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/carlosbenavides123/DevJobs/MainServer/dbconf"
	"github.com/carlosbenavides123/DevJobs/MainServer/mvc/utils"
	uuid "github.com/nu7hatch/gouuid"
)

func GetJobs(jobIdx string) (*JobResponse, *utils.ApplicationError) {

	fmt.Println(jobIdx, "GET JOBS TIMESTAMP")

	var index = 0

	db := dbconf.DbConn()
	defer db.Close()

	stmt, dbPrepareErr := db.Prepare(`select j.job_uuid, j.job_title, j.job_link, j.job_location, j.job_posted, 
						j.job_idx, c.company_name, c.company_cloudinary, l.job_level
						from companies c 
						inner join jobs_pivot jp on jp.company_uuid = c.company_uuid 
						inner join jobs j on jp.job_uuid = j.job_uuid 
						inner join levels l on j.experience_level = l.id
						WHERE j.job_idx <= ? AND active = 1
						ORDER BY j.job_idx DESC
						LIMIT 21`)

	if dbPrepareErr != nil {
		panic(dbPrepareErr.Error())
	}

	res, dbErr := stmt.Query(jobIdx)

	if dbErr != nil {
		panic(dbErr.Error())
	}

	job := []*Job{}

	cursor := &Cursor{}
	jobResponse := &JobResponse{}

	for res.Next() {
		var JobUUID, JobTitle, JobLink, JobLocation string
		var JobPosted, JobIdx int64
		var CompanyName, Cloudinary, JobLevel string

		scanErr := res.Scan(&JobUUID, &JobTitle, &JobLink, &JobLocation, &JobPosted, &JobIdx, &CompanyName, &Cloudinary, &JobLevel)
		if scanErr != nil {
			panic(scanErr.Error())
		}

		fmt.Println(JobIdx, index)

		if index == 20 {
			cursor.Cursor = JobIdx
			break
		} else {
			jobRow := &Job{}
			jobRow.JobUUID = JobUUID
			jobRow.JobTitle = JobTitle
			jobRow.JobLink = JobLink
			jobRow.JobLocation = JobLocation
			jobRow.JobPosted = JobPosted
			jobRow.JobIdx = JobIdx
			jobRow.CompanyName = CompanyName
			jobRow.Cloudinary = Cloudinary
			jobRow.JobLevel = JobLevel
			job = append(job, jobRow)
		}

		index++
	}

	if dbErr == nil {
		jobResponse.Job = job
		jobResponse.Cursor = cursor
		return jobResponse, nil
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

	stmt, dbPrepareErr := db.Prepare(`select j.job_uuid, j.job_title, j.job_link, j.job_location, j.job_posted, j.job_found, c.company_name, c.company_cloudinary, l.job_level
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
		var JobPosted, JobFound int64
		var CompanyName, Cloudinary, JobLevel string

		scanErr := res.Scan(&JobUUID, &JobTitle, &JobLink, &JobLocation, &JobPosted, &JobFound, &CompanyName, &Cloudinary, &JobLevel)

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

func GetJobsByLocation(location string, jobIdx string) (*JobResponse, *utils.ApplicationError) {
	var index = 0

	location = "%" + location + "%"

	db := dbconf.DbConn()
	defer db.Close()

	stmt, dbPrepareErr := db.Prepare(`select j.job_uuid, j.job_title, j.job_link, j.job_location, j.job_posted, 
						j.job_idx, c.company_name, c.company_cloudinary, l.job_level
						from companies c 
						inner join jobs_pivot jp on jp.company_uuid = c.company_uuid 
						inner join jobs j on jp.job_uuid = j.job_uuid 
						inner join levels l on j.experience_level = l.id
						WHERE j.job_location LIKE ?
						AND j.job_idx <= ? 
						AND active = 1
						ORDER BY j.job_idx DESC
						LIMIT 21`)
	if dbPrepareErr != nil {
		panic(dbPrepareErr.Error())
	}

	res, queryErr := stmt.Query(location, jobIdx)

	if queryErr != nil {
		panic(queryErr.Error())
	}

	job := []*Job{}

	cursor := &Cursor{}
	jobResponse := &JobResponse{}

	for res.Next() {
		var JobUUID, JobTitle, JobLink, JobLocation string
		var JobPosted, JobIdx int64
		var CompanyName, Cloudinary, JobLevel string

		scanErr := res.Scan(&JobUUID, &JobTitle, &JobLink, &JobLocation, &JobPosted, &JobIdx, &CompanyName, &Cloudinary, &JobLevel)
		if scanErr != nil {
			panic(scanErr.Error())
		}

		fmt.Println(JobIdx, index)

		if index == 20 {
			cursor.Cursor = JobIdx
			break
		} else {
			jobRow := &Job{}
			jobRow.JobUUID = JobUUID
			jobRow.JobTitle = JobTitle
			jobRow.JobLink = JobLink
			jobRow.JobLocation = JobLocation
			jobRow.JobPosted = JobPosted
			jobRow.JobIdx = JobIdx
			jobRow.CompanyName = CompanyName
			jobRow.Cloudinary = Cloudinary
			jobRow.JobLevel = JobLevel
			job = append(job, jobRow)
		}

		index++
	}

	if index != 20 {
		cursor.Cursor = 0
	}

	jobResponse.Cursor = cursor
	jobResponse.Job = job
	return jobResponse, nil

}

func GetJobsByExperience(experience string, cursor string) (*JobResponse, *utils.ApplicationError) {
	fmt.Println(experience)
	if experience == "Intern" {
		return queryIntern(cursor)
	} else if experience == "Entry" {
		return queryEntry(cursor)
	} else if experience == "Mid" {

	} else if experience == "Senior" {
		return querySenior(cursor)
	}
	return nil, nil
}

func queryEntry(cursor string) (*JobResponse, *utils.ApplicationError) {
	db := dbconf.DbConn()
	defer db.Close()

	stmt, dbPrepareErr := db.Prepare(`SELECT 
									ej.job_uuid, ej.job_title, ej.job_link, ej.job_location, ej.job_posted,  ej.job_idx, 
									c.company_name, c.company_cloudinary 
									FROM companies c  
									INNER JOIN jobs_pivot jp ON
									jp.company_uuid = c.company_uuid  
									INNER JOIN entry_jobs ej ON
									jp.job_uuid = ej.job_uuid 
									WHERE ej.job_idx <= ?  
									AND active = 1 
									ORDER BY ej.job_idx DESC 
									LIMIT 21;`)
	if dbPrepareErr != nil {
		panic(dbPrepareErr.Error())
	}
	res, queryErr := stmt.Query(cursor)
	if queryErr != nil {
		panic(queryErr.Error())
	}
	jobResponse := experienceTableScan(res)
	return jobResponse, nil
}

func queryIntern(cursor string) (*JobResponse, *utils.ApplicationError) {
	db := dbconf.DbConn()
	defer db.Close()

	stmt, dbPrepareErr := db.Prepare(`SELECT 
									ij.job_uuid, ij.job_title, ij.job_link, ij.job_location, ij.job_posted,  ij.job_idx, 
									c.company_name, c.company_cloudinary 
									FROM companies c  
									INNER JOIN jobs_pivot jp ON
									jp.company_uuid = c.company_uuid  
									INNER JOIN intern_jobs ij ON
									jp.job_uuid = ij.job_uuid 
									WHERE ij.job_idx <= ?  
									AND active = 1 
									ORDER BY ij.job_idx DESC 
									LIMIT 21;`)
	if dbPrepareErr != nil {
		panic(dbPrepareErr.Error())
	}
	res, queryErr := stmt.Query(cursor)
	if queryErr != nil {
		panic(queryErr.Error())
	}
	jobResponse := experienceTableScan(res)
	return jobResponse, nil
}

func querySenior(cursor string) (*JobResponse, *utils.ApplicationError) {
	db := dbconf.DbConn()
	defer db.Close()

	stmt, dbPrepareErr := db.Prepare(`SELECT 
									sj.job_uuid, sj.job_title, sj.job_link, sj.job_location, sj.job_posted,  sj.job_idx, 
									c.company_name, c.company_cloudinary 
									FROM companies c  
									INNER JOIN jobs_pivot jp ON
									jp.company_uuid = c.company_uuid  
									INNER JOIN senior_jobs sj ON
									jp.job_uuid = sj.job_uuid 
									WHERE sj.job_idx <= ?  
									AND active = 1 
									ORDER BY sj.job_idx DESC 
									LIMIT 21;`)
	if dbPrepareErr != nil {
		panic(dbPrepareErr.Error())
	}
	res, queryErr := stmt.Query(cursor)
	if queryErr != nil {
		panic(queryErr.Error())
	}
	jobResponse := experienceTableScan(res)
	return jobResponse, nil
}

func experienceTableScan(res *sql.Rows) *JobResponse {
	var index = 0
	job := []*Job{}
	cursor := &Cursor{}
	jobResponse := &JobResponse{}

	for res.Next() {
		var JobUUID, JobTitle, JobLink, JobLocation string
		var JobPosted, JobIdx int64
		var CompanyName, Cloudinary string

		scanErr := res.Scan(&JobUUID, &JobTitle, &JobLink, &JobLocation, &JobPosted, &JobIdx, &CompanyName, &Cloudinary)
		if scanErr != nil {
			panic(scanErr.Error())
		}
		if index == 20 {
			cursor.Cursor = JobIdx
		} else {
			jobRow := &Job{}
			jobRow.JobUUID = JobUUID
			jobRow.JobTitle = JobTitle
			jobRow.JobLink = JobLink
			jobRow.JobLocation = JobLocation
			jobRow.JobPosted = JobPosted
			jobRow.JobIdx = JobIdx
			jobRow.CompanyName = CompanyName
			jobRow.Cloudinary = Cloudinary
			job = append(job, jobRow)
		}
		index++
	}
	jobResponse.Cursor = cursor
	jobResponse.Job = job
	return jobResponse
}
