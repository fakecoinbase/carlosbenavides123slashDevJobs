package services

import (
	"fmt"
	"net/http"

	location "github.com/carlosbenavides123/DevJobs/MainServer/mvc/domain/Location"
	"github.com/carlosbenavides123/DevJobs/MainServer/mvc/utils"
	"github.com/gorilla/mux"
)

func GetCompaniesByLocation(r *http.Request) ([]*location.Location, *utils.ApplicationError) {
	params := mux.Vars(r)
	loc := params["location"]
	if loc == "" {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Location is required!"),
			StatusCode: http.StatusNotFound,
			Code:       "No Jobs available.",
		}
	}
	return location.GetCompaniesByLocation(loc)
}

func GetLocationsByCompany(r *http.Request) ([]*location.Location, *utils.ApplicationError) {
	params := mux.Vars(r)
	comp := params["company"]
	if comp == "" {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Company name is required!"),
			StatusCode: http.StatusNotFound,
			Code:       "No Jobs available.",
		}
	}
	return location.GetLocationsByCompany(comp)
}
