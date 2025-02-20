package main

import (
	"fmt"
	"log"
	"net/http"
	"tmu-webring-go/internal/account"
	"tmu-webring-go/internal/middleware"
	"tmu-webring-go/internal/submission"
	"tmu-webring-go/internal/website"
)

func main() {
	router := http.NewServeMux()

	authRouter := http.NewServeMux()
	/*
	 * GET: return account information and settings
	 * PATCH: update account information and settings
	 * POST: create account
	 * DELETE: delete account
	**/
	router.HandleFunc("POST /account", account.CreateAccount)
	authRouter.HandleFunc("GET /account/{username}", account.GetAccount)
	authRouter.HandleFunc("PATCH /account/{username}", account.UpdateAccount)
	authRouter.HandleFunc("DELETE /account/{username}", account.DeleteAccount)

	/*
	 * GET: all approved websites
	**/
	router.HandleFunc("GET /websites", website.GetWebsites)

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

	server := http.Server{
		Addr:    ":8021",
		Handler: stack(router),
	}
	fmt.Println("Starting GO API service...")
	fmt.Println(`
 ______     ______        ______     ______   __    
/\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \   
\ \ \__ \  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \  
 \ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\ 
  \/_____/   \/_____/      \/_/\/_/   \/_/     \/_/ `)

	log.Fatal(server.ListenAndServe())
}
