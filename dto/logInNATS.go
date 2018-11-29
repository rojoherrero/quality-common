package dto

import (
	"encoding/json"
)

type LogInRequestNATS struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LogInResponseNATS struct {
	UsernameOK bool `json:"usernameOK"`
	PasswordOK *bool `json:"passwordOK"`
	Roles []string `json:"roles"`
	Deps []string `json:"deps"`
	FullName *string `json:"fullName"`
}

func UnmarshalLogInResponseNATS(data []byte) (LogInResponseNATS, error) {
	var r LogInResponseNATS
	e := json.Unmarshal(data, &r)
	return r, e
}

func (r *LogInResponseNATS) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalLogInRequestNATS(data []byte) (LogInRequestNATS, error) {
	var r LogInRequestNATS
	e := json.Unmarshal(data, &r)
	return r, e
}

func (r *LogInRequestNATS) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
