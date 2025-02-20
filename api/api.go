package api

import (
	"encoding/json"
	"net/http"
)

type CreateAccountParams struct {
	Username string
	Email    string
	Password string
}

type AccountParams struct {
	Username string
}

type SubmissionParams struct {
	Name      string
	URL       string
	Year      int8
	Skills    [2]string
	Employers [2]string
}

type SubmissionResponse struct {
	Name      string
	URL       string
	Year      int8
	Skills    [2]string
	Employers [2]string
}

type Error struct {
	Code    int
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}

	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}

	UnauthorizedErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Unauthorized.", http.StatusUnauthorized)
	}
)
