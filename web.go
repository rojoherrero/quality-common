package common

import (
	"net/http"
)

const (
	contentType = "Content-Type"
	jsonUTF8    = "application/json; charset=UTF-8"
)

func JSON(w http.ResponseWriter, status int, data []byte) error {
	w.Header().Set(contentType, jsonUTF8)
	w.WriteHeader(status)
	_, e := w.Write(data)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return e
	}
	return nil
}
