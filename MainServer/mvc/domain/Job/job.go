package job

type Job struct {
	UUID        string `json:"uuid"`
	CompanyName string `json:"company_name"`
	JobLink     string `json:"job_link"`
	JobPosted   int64  `json:"job_posted"`
	JobFound    int64  `json:"job_found"`
}
