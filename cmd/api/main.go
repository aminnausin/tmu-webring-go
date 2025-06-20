package main

import (
	"database/sql"
	"log"
	"net/http"
	"tmu-webring-go/internal/env"
	"tmu-webring-go/internal/mailer"
	"tmu-webring-go/internal/middleware"
	"tmu-webring-go/internal/storage"
	"tmu-webring-go/internal/submission"
	"tmu-webring-go/internal/web"

	_ "modernc.org/sqlite"
	// _ "github.com/mattn/go-sqlite3" // <-- THIS is required
)

func main() {
	env.ReadEnvFile()

	db, err := sql.Open("sqlite", "file:test.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	brevoMailer, err := mailer.NewBrevoClient(env.GetStringEnv("MAIL_USERNAME", ""), env.GetStringEnv("MAIL_PASSWORD", ""), "no-reply@nausin.me")
	if err != nil {
		log.Fatal(err)
	}

	storage := storage.NewStorage(db)

	// userService := user.NewUserService(storage, brevoMailer)
	app := &application{storage: storage, mailer: brevoMailer}

	router := http.NewServeMux()
	router.HandleFunc("POST /api/user", app.registerUserHandler)

	authRouter := http.NewServeMux()

	// router.HandleFunc("GET /", nil)

	/*
	 * GET: return account information and settings
	 * PATCH: update account information and settings
	 * POST: create account
	 * DELETE: delete account
	**/
	// router.HandleFunc("POST /account", user.CreateAccount)
	// authRouter.HandleFunc("GET /account/{username}", user.GetAccount)
	// authRouter.HandleFunc("PATCH /account/{username}", user.UpdateAccount)
	// authRouter.HandleFunc("DELETE /account/{username}", user.DeleteAccount)

	/*
	 * GET: all approved websites
	**/
	// router.HandleFunc("GET /websites", website.GetWebsites)
	web.RegisterStaticRoutes(router)
	/*
	 * GET: get submissions
	 * PATCH: approve submission
	 * POST: create submission
	 * DELETE: delete submission
	**/
	router.HandleFunc("POST /submissions", submission.CreateSubmission)
	authRouter.HandleFunc("GET /submissions", submission.NewHandler().GetSubmissions)
	authRouter.HandleFunc("PATCH /submissions/{id}", submission.NewHandler().UpdateSubmission)
	authRouter.HandleFunc("DELETE /submissions/{id}", submission.NewHandler().DeleteSubmission)

	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))

	// router.Handle("/", middleware.EnsureAdmin(adminRouter))
	stack := middleware.CreateStack(
		middleware.Logging,
	)

	log.Fatal(app.run(stack(router)))
}
