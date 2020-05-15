package services

import (
	"net/http"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

	cms "github.com/carlosbenavides123/DevJobs/MainServer/mvc/domain/Cms"
	"github.com/carlosbenavides123/DevJobs/MainServer/pb/company/companypb"
)

func GetCmsHomeData(r *http.Request) (companyResponse *companypb.CompanyResponse) {
	return cms.GetCmsHomeData()
}

func GetCmsCompanyData(p *kafka.Producer, c *kafka.Consumer, r *http.Request) *cms.CompanyCms {
	company := r.URL.Query().Get("company")
	return cms.GetCmsCompanyData(p, c, company)
}
