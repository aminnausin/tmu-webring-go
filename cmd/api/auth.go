package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"tmu-webring-go/internal/mailer"
	"tmu-webring-go/internal/storage"

	"github.com/google/uuid"
)

type RegisterUserRequest struct {
	Name         string `json:"name" validate:"required,max=100"`
	Email        string `json:"email" validate:"required,email,max=255"`
	Password     string `json:"password" validate:"required,min=3,max=72"`
	Confirmation string `json:"password_confirmation" validate:"required,min=3,max=72"`
}

type UserWithToken struct {
	*storage.User
	Token string `json:"token"`
}

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterUserRequest
	if err := readJSON(w, r, &req); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(req); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if req.Password != req.Confirmation {
		app.badRequestResponse(w, r, errors.New("password and confirmation do not match"))
		return
	}

	user := &storage.User{
		Name:  req.Name,
		Email: req.Email,
	}

	// hash the user password
	if err := user.Password.Set(req.Password); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	plainToken := uuid.New().String()

	// hash the token for storage but keep the plain token for email
	// hash := sha256.Sum256([]byte(plainToken))
	// hashToken := hex.EncodeToString(hash[:])
	err := app.storage.Users.CreateAndInvite(ctx, user)
	if err != nil {
		log.Println(err)

		switch err {
		case storage.ErrDuplicateEmail:
			app.badRequestResponse(w, r, err)
		case storage.ErrDuplicateUsername:
			app.badRequestResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	userWithToken := UserWithToken{
		User:  user,
		Token: plainToken,
	}
	activationURL := fmt.Sprintf("%s/confirm/%s", "127.0.0.1:8021", plainToken)

	isProdEnv := false
	vars := struct {
		Name          string
		ActivationURL string
	}{
		Name:          user.Name,
		ActivationURL: activationURL,
	}

	// send mail
	status, err := app.mailer.Send(mailer.UserWelcomeTemplate, user.Name, user.Email, vars, !isProdEnv)
	if err != nil {
		// app.logger.Errorw("error sending welcome email", "error", err)

		// rollback user creation if email fails (SAGA pattern)
		// if err := app.storage.Users.Delete(ctx, user.ID); err != nil {
		// app.logger.Errorw("error deleting user", "error", err)
		// }
		log.Println(err)
		app.internalServerError(w, r, err)
		return
	}

	log.Printf("Email sent with status code: %d", status)
	// app.logger.Infow("Email sent", "status code", status)

	if err := app.jsonResponse(w, http.StatusCreated, userWithToken); err != nil {
		app.internalServerError(w, r, err)
	}

	w.WriteHeader(http.StatusCreated)
}
