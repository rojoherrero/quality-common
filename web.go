package common

import (
	"encoding/json"
	"net/http"
)

const (
	contentType = "Content-Type"
	jsonUTF8    = "application/json; charset=UTF-8"
)

func JSON(res http.ResponseWriter, status int, data []byte) error {
	body, e := json.Marshal(data)
	if e != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return e
	}

	res.Header().Set(contentType, jsonUTF8)
	res.WriteHeader(status)
	_, e = res.Write(body)
	if e != nil {
		return e
	}
	return nil
}
