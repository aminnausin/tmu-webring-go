package submission

import (
	"log"
	"net/http"
	"tmu-webring-go/internal/middleware"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetSubmissions(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.AuthUserID).(string)

	if !ok {
		log.Println("invalid user ID")
		w.WriteHeader(http.StatusBadRequest)
	}

	log.Println("X for ", userID)

	w.Write([]byte("x done for " + userID))
}

func (h *Handler) UpdateSubmission(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.AuthUserID).(string)

	if !ok {
		log.Println("invalid user ID")
		w.WriteHeader(http.StatusBadRequest)
	}

	log.Println("X for ", userID)

	w.Write([]byte("x done for " + userID))
}

func (h *Handler) DeleteSubmission(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.AuthUserID).(string)

	if !ok {
		log.Println("invalid user ID")
		w.WriteHeader(http.StatusBadRequest)
	}

	log.Println("X for ", userID)

	w.Write([]byte("x done for " + userID))
}

func CreateSubmission(w http.ResponseWriter, r *http.Request) {

}
