package location

import (
	"fmt"
	"net/http"

	"github.com/carlosbenavides123/DevJobs/MainServer/dbconf"
	"github.com/carlosbenavides123/DevJobs/MainServer/mvc/utils"
)

func GetCompaniesByLocation(location string) ([]*Location, *utils.ApplicationError) {
	db := dbconf.DbConn()
	defer db.Close()

	stmt, err := db.Prepare(`select l.location, l.company_name 
						from locations l
						WHERE l.location=?`)
	if err != nil {
		panic(err.Error())
	}

	res, queryErr := stmt.Query(location)

	if queryErr != nil {
		panic(error(queryErr))
	}

	locations := []*Location{}
	for res.Next() {
		var LocationName, CompanyName string
		scanErr := res.Scan(&LocationName, &CompanyName)

		if scanErr != nil {
			panic(scanErr.Error())
		}
		location := &Location{}
		location.Location = LocationName
		location.CompanyName = CompanyName
		locations = append(locations, location)
	}
	if len(locations) == 0 {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Companies were not found"),
			StatusCode: http.StatusNotFound,
			Code:       "not found",
		}
	}
	return locations, nil
}
