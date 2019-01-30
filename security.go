package common

import (
	"encoding/json"
)

// LogInRequest holds user credentials
type LogInRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// UnmarshalLogInRequest Unmarshal an array of bytes to the struct
func UnmarshalLogInRequest(data []byte) (LogInRequest, error) {
	var r LogInRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal the struct to an byte array
func (r *LogInRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// LogInResponse holds the resulto of the authetication
type LogInResponse struct {
	// ResponseStatus save status as
	ResponseStatus int      `json:"responseStatus,omitempty"`
	Username       string   `json:"username,omitempty"`
	Fullname       string   `json:"fullname,omitempty"`
	Department     string   `json:"department,omitempty"`
	Roles          []string `json:"roles,omitempty"`
}

// UnmarshalLogInResponse Unmarshal an array of bytes to the struct
func UnmarshalLogInResponse(data []byte) (LogInResponse, error) {
	var r LogInResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal the struct to an byte array
func (r *LogInResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

const (
	// LogInService constant that holds the value os the nats route
	LogInService string = "logIn.security.service"
)
