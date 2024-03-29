package main

import (
	"log"
	"matthewhope/atm-system-go/router"
	"matthewhope/atm-system-go/ui"
	"net/http"
	"regexp"
)

func main() {
	mux := http.NewServeMux()
	serveStaticFiles(mux)
	r := router.New()
	r.AddHandler(regexp.MustCompile(`/$`), ui.NewIndexHandler())
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
