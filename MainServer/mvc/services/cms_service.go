package services

import (
	"net/http"

	cms "github.com/carlosbenavides123/DevJobs/MainServer/mvc/domain/Cms"
	"github.com/carlosbenavides123/DevJobs/MainServer/pb/company/companypb"
)

func GetCmsHomeData(r *http.Request) (companyResponse *companypb.CompanyResponse) {
	return cms.GetCmsHomeData()
}

func GetCmsCompanyData(r *http.Request) {
	company := r.URL.Query().Get("company")
	cms.GetCmsCompanyData(company)
}
