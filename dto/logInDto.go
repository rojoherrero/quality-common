package dto

import (
	"encoding/json"
)

type LogInRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LogInResponseDTO struct {
	UsernameOK bool `json:"usernameOK"`
	PasswordOK *bool `json:"passwordOK"`
	Roles []string `json:"roles"`
	Deps []string `json:"deps"`
	FullName *string `json:"fullName"`
}

func UnmarshalLogInResponseDTO(data []byte) (LogInResponseDTO, error) {
	var r LogInResponseDTO
	e := json.Unmarshal(data, &r)
	return r, e
}

func (r *LogInResponseDTO) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalLogInRequestDTO(data []byte) (LogInRequestDTO, error) {
	var r LogInRequestDTO
	e := json.Unmarshal(data, &r)
	return r, e
}

func (r *LogInRequestDTO) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
