package common

import (
	"encoding/json"
	"net/http"
)

const (
	contentType = "Content-Type"
	jsonUTF8    = "application/json; charset=UTF-8"
)

func JSON(w http.ResponseWriter, status int, data interface{}) error {
	body, e := json.Marshal(data)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return e
	}

	w.Header().Set(contentType, jsonUTF8)
	w.WriteHeader(status)
	_, e = w.Write(body)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return e
	}
	return nil
}
