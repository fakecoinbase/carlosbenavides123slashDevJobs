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

	insForm1, err := db.Prepare(`INSERT INTO 
				jobs_pivot(job_uuid, company_uuid)
				VALUES(?, ?)`)
	if err != nil {
		panic(err.Error())
	}
	insForm1.Exec(job.JobID, job.Company_UUID)

	insForm2, err := db.Prepare(`INSERT INTO 
							jobs(job_uuid, company_uuid, job_title, job_link,
								job_posted, job_found, active, experience_level)
							VALUES(?,?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		panic(err.Error())
	}
	splitString := strings.Split(job.JobID, "_%_")
	jobTitle := strings.ReplaceAll(splitString[1], "%", " ")

	var level = 1

	if job.Entry {
		level = 1
	} else if job.Entry {
		level = 2
	} else if job.Mid {
		level = 3
	} else if job.Senior {
		level = 4
	} else {
		level = 5
	}
	insForm2.Exec(job.JobID, job.Company_UUID, jobTitle, job.JobLink, job.JobPosted, msg.Timestamp.Unix(), job.Active, level)
	defer db.Close()

}
