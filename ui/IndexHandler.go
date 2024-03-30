package ui

import (
	"fmt"
	"log"
	"net/http"
)

type IndexViewModel struct {
	Title string
}

type IndexHandler struct {
	viewData IndexViewModel
}

func NewIndexHandler() *IndexHandler {
	ret := new(IndexHandler)
	ret.viewData = IndexViewModel{Title: "Home"}
	return ret
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("Session")
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Printf("%+v\n", c)
	}
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *IndexHandler) get(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)
	err := tmpl.ExecuteTemplate(w, "Index", h.viewData)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
