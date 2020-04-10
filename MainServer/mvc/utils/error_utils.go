package utils

type ApplicationError struct {
	Message    string `json:"Message"`
	StatusCode int    `json:"Status"`
	Code       string `json:"Code"`
}
