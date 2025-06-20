package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tmu-webring-go/internal/mailer"
	"tmu-webring-go/internal/storage"
)

type application struct {
	storage storage.Storage
	mailer  mailer.Client
}

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

func (app *application) run(mux http.Handler) error {
	// Docs
	http := http.Server{
		Addr:    ":8021",
		Handler: mux,
	}
	fmt.Println("Starting GO API service on port 8021...")
	fmt.Println(`
 ______     ______        ______     ______   __    
/\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \   
\ \ \__ \  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \  
 \ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\ 
  \/_____/   \/_____/      \/_/\/_/   \/_/     \/_/ `)

	if err := http.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	return nil
}
