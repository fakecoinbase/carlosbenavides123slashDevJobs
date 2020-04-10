package consumergroups

import (
	"log"

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

	insForm, err := db.Prepare(`INSERT INTO 
							jobs(UUID, CompanyName, JobLink, JobPosted,
							JobFound, Active)
							VALUES(?,?, ?, ?, ?, ?)`)
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(job.JobID, job.CompanyName, job.JobLink, job.JobPosted, msg.Timestamp.Unix(), job.Active)
	defer db.Close()

}
