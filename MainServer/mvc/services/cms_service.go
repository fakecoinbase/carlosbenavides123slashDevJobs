package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/carlosbenavides123/DevJobs/MainServer/mvc/utils"

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

func UpdateCompanyCMSData(p *kafka.Producer, c *kafka.Consumer, r *http.Request) (*utils.ApplicationSuccess, *utils.ApplicationError) {
	decoder := json.NewDecoder(r.Body)
	var updateCompanyDetails *cms.CompanyCms
	decodeErr := decoder.Decode(&updateCompanyDetails)

	if decodeErr != nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Missing one or more paramters."),
			StatusCode: http.StatusBadRequest,
			Code:       "Malformed request",
		}
	}

	if updateCompanyDetails.CompanyUUID == "" {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Missing Company UUID! (Required)"),
			StatusCode: http.StatusBadRequest,
			Code:       "Malformed request",
		}
	}

	if updateCompanyDetails.CompanyName == "" {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Missing Company Name! (Required)"),
			StatusCode: http.StatusBadRequest,
			Code:       "Malformed request",
		}
	}

	if updateCompanyDetails.CompanyWebsite == "" {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Missing Company Website! (Required)"),
			StatusCode: http.StatusBadRequest,
			Code:       "Malformed request",
		}
	}

	if updateCompanyDetails.GreenHouse == false && updateCompanyDetails.Lever == false && updateCompanyDetails.Other == false {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Atleast one radio button has to be chosen! (Required)"),
			StatusCode: http.StatusBadRequest,
			Code:       "Malformed request",
		}
	}

	if strings.Contains(updateCompanyDetails.WantedDepartments, ", ") {
		updateCompanyDetails.WantedDepartments = strings.ReplaceAll(updateCompanyDetails.WantedDepartments, ", ", ",")
	}

	return cms.UpdateCompanyCMSData(updateCompanyDetails, p)
}
