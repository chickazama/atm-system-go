package api

import (
	"encoding/json"
	"io"
	"log"
	"matthewhope/atm-system-go/models"
	"matthewhope/atm-system-go/repo"
	"net/http"
)

type UsersHandler struct {
	rp repo.IRepository
}

func NewUsersHandler(r repo.IRepository) *UsersHandler {
	ret := new(UsersHandler)
	ret.rp = r
	return ret
}

func (h *UsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// CREATE
	case http.MethodPost:
		h.post(w, r)
	}
}

func (h *UsersHandler) post(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	switch contentType {
	case "application/json":
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		var u models.User
		err = json.Unmarshal(buf, &u)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		ret, err := h.rp.CreateUser(u)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(&ret)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "unsupported media type", http.StatusUnsupportedMediaType)
	}
}
