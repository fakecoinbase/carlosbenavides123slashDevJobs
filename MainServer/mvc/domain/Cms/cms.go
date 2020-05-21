package cms

type CompanyCms struct {
	CompanyUUID       string `json:"company_uuid"`
	CompanyName       string `json:"company_name"`
	CompanyWebsite    string `json:"company_website"`
	WantedDepartments string `json:"wanted_departments"`
	WantedLocations   string `json:"wanted_locations"`
	GreenHouse        bool   `json:"greenhouse"`
	Lever             bool   `json:"lever"`
	Other             bool   `json:"other"`
}
