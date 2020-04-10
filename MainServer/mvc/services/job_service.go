package services

import (
	job "github.com/carlosbenavides123/DevJobs/MainServer/mvc/domain/Job"
	"github.com/carlosbenavides123/DevJobs/MainServer/mvc/utils"
)

func GetJobs() ([]*job.Job, *utils.ApplicationError) {
	return job.GetJobs()
}
