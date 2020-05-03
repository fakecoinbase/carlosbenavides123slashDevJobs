package consumergroups

import (
	"log"
	"strings"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

	"github.com/carlosbenavides123/DevJobs/MainServer/dbconf"
	"github.com/carlosbenavides123/DevJobs/MainServer/jobpb"
	"github.com/golang/protobuf/proto"
)

func Addnewjob(msg *kafka.Message) {

	job := &jobpb.Job{}
	if err := proto.Unmarshal(msg.Value, job); err != nil {
		log.Fatalln("Failed to parse Job:", err)
	}
	db := dbconf.DbConn()
	defer db.Close()

	insForm1, err := db.Prepare(`INSERT INTO 
				jobs_pivot(job_uuid, company_uuid)
				VALUES(?, ?)`)
	if err != nil {
		panic(err.Error())
	}
	insForm1.Exec(job.JobUUID, job.CompanyUUID)

	insForm2, err := db.Prepare(`INSERT INTO 
							jobs(job_uuid, company_uuid, job_title, job_link, job_location,
								job_posted, job_found, active, experience_level)
							VALUES(?,?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		panic(err.Error())
	}
	splitString := strings.Split(job.JobUUID, "_%_")
	jobTitle := strings.ReplaceAll(splitString[1], "%", " ")
	jobLocation := strings.ReplaceAll(splitString[2], "%", " ")

	if job.JobPosted == 0 {
		job.JobPosted = msg.Timestamp.Unix()
	}

	insForm2.Exec(job.JobUUID, job.CompanyUUID, jobTitle, job.JobLink, jobLocation, job.JobPosted, msg.Timestamp.Unix(), job.Active, job.ExperienceLevel)

	if job.ExperienceLevel == 1 {
		replicateInternJobData(job, jobTitle, jobLocation)
	} else if job.ExperienceLevel == 2 {
		replicateEntryJobData(job, jobTitle, jobLocation)
	} else if job.ExperienceLevel == 3 {
		replicateMidJobs(job, jobTitle, jobLocation)
	} else if job.ExperienceLevel == 4 {
		replicateSeniorJobs(job, jobTitle, jobLocation)
	}
}

func replicateInternJobData(job *jobpb.Job, jobtitle string, jobLocation string) {
	db := dbconf.DbConn()
	defer db.Close()

	stmt, prepareErr := db.Prepare(`INSERT INTO intern_jobs(job_uuid, company_uuid, job_title,
									job_link, job_location, job_posted,
									active) VALUES
									(?, ?, ?, ?, ?, ?, ?)`)
	if prepareErr != nil {
		panic(prepareErr.Error())
	}
	stmt.Exec(job.JobUUID, job.CompanyUUID, jobtitle, job.JobLink, jobLocation, job.JobPosted, job.Active)
}

func replicateEntryJobData(job *jobpb.Job, jobtitle string, jobLocation string) {
	db := dbconf.DbConn()
	defer db.Close()

	stmt, prepareErr := db.Prepare(`INSERT INTO entry_jobs(job_uuid, company_uuid, job_title,
									job_link, job_location, job_posted,
									active) VALUES
									(?, ?, ?, ?, ?, ?, ?)`)
	if prepareErr != nil {
		panic(prepareErr.Error())
	}
	stmt.Exec(job.JobUUID, job.CompanyUUID, jobtitle, job.JobLink, jobLocation, job.JobPosted, job.Active)
}

func replicateMidJobs(job *jobpb.Job, jobtitle string, jobLocation string) {
	db := dbconf.DbConn()
	defer db.Close()

	stmt, prepareErr := db.Prepare(`INSERT INTO mid_jobs(job_uuid, company_uuid, job_title,
									job_link, job_location, job_posted,
									active) VALUES
									(?, ?, ?, ?, ?, ?, ?)`)
	if prepareErr != nil {
		panic(prepareErr.Error())
	}
	stmt.Exec(job.JobUUID, job.CompanyUUID, jobtitle, job.JobLink, jobLocation, job.JobPosted, job.Active)
}

func replicateSeniorJobs(job *jobpb.Job, jobtitle string, jobLocation string) {
	db := dbconf.DbConn()
	defer db.Close()

	stmt, prepareErr := db.Prepare(`INSERT INTO senior_jobs(job_uuid, company_uuid, job_title,
									job_link, job_location, job_posted,
									active) VALUES
									(?, ?, ?, ?, ?, ?, ?)`)
	if prepareErr != nil {
		panic(prepareErr.Error())
	}
	stmt.Exec(job.JobUUID, job.CompanyUUID, jobtitle, job.JobLink, jobLocation, job.JobPosted, job.Active)
}
