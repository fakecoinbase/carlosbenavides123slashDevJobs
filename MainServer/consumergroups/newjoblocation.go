package consumergroups

import (
	"log"

	"github.com/carlosbenavides123/DevJobs/MainServer/dbconf"
	"github.com/carlosbenavides123/DevJobs/MainServer/pb/locationpb"
	"github.com/golang/protobuf/proto"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func AddNewJobLocation(msg *kafka.Message) {
	locationpb := &locationpb.Location{}
	if err := proto.Unmarshal(msg.Value, locationpb); err != nil {
		log.Fatalln("Failed to parse new location message:", err)
	}
	db := dbconf.DbConn()
	defer db.Close()

	insertJobLoc, dbPrepErr := db.Prepare(`INSERT INTO 
								locations(location, company_name)
								VALUES(?,?)`)

	if dbPrepErr != nil {
		panic(dbPrepErr.Error())
	}

	insertJobLoc.Exec(locationpb.Location, locationpb.CompanyName)
}
