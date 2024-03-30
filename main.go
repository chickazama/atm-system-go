package main

import (
	"log"
	"matthewhope/atm-system-go/api"
	"matthewhope/atm-system-go/auth"
	"matthewhope/atm-system-go/repo"
	"matthewhope/atm-system-go/router"
	"matthewhope/atm-system-go/ui"
	"net/http"
	"regexp"
)

func main() {
	mux := http.NewServeMux()
	serveStaticFiles(mux)
	r := router.New()
	rp := repo.NewSQLiteRepository()
	r.AddHandler(regexp.MustCompile(`/$`), ui.NewIndexHandler())
	addAPIHandlers(r, rp)
	mux.Handle("/", r)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func serveStaticFiles(mux *http.ServeMux) {
	fsRoot := http.Dir("./static/")
	fs := http.FileServer(fsRoot)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
}

func addAPIHandlers(r *router.Router, rp repo.IRepository) {
	r.AddHandler(regexp.MustCompile(`^/api/users$`), api.NewUsersHandler(rp))
	r.AddHandler(regexp.MustCompile(`^/api/accounts$`), api.NewAccountsHandler(rp))
	r.AddHandler(regexp.MustCompile(`^/auth/github$`), auth.NewGithubHandler())
	r.AddHandler(regexp.MustCompile(`^/auth/callback$`), auth.NewCallbackHandler())
}
