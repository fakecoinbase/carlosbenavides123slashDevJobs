package job

type Job struct {
	JobUUID     string `json:"job_uuid"`
	JobTitle    string `json:"job_title"`
	JobLink     string `json:"job_link"`
	JobLocation string `json:"job_location"`
	JobPosted   int64  `json:"job_posted"`
	JobFound    int64  `json:"job_found"`
	JobIdx      int64  `json:"job_idx"`
	CompanyName string `json:"company_name"`
	Cloudinary  string `json:"cloudinary"`
	JobLevel    string `json:"level"`
}

type Cursor struct {
	Cursor int64 `json:"next_cursor"`
}

type JobResponse struct {
	Job    []*Job
	Cursor *Cursor
}

type NewCompany struct {
	CompanyName    string `json:"company_name"`
	CompanyWebsite string `json:"company_website"`
	Cloudinary     string `json:"cloudinary"`
}

type Company struct {
	CompanyName string `json:"company_name"`
	CompanyUUID string `json:"company_uuid"`
}
