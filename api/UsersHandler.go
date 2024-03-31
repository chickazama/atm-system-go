package api

import (
	"encoding/json"
	"io"
	"log"
	"matthewhope/atm-system-go/models"
	"matthewhope/atm-system-go/repo"
	"matthewhope/atm-system-go/transport"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
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
	// READ
	case http.MethodGet:
		h.get(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// CREATE (HTTP POST)
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
		var dto transport.CreateUserDTO
		err = json.Unmarshal(buf, &dto)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		ctime := time.Now().UTC().UnixMilli()
		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		u := models.User{
			CreatedAt:         ctime,
			EmailAddress:      dto.EmailAddress,
			EncryptedPassword: string(encryptedPassword),
			FirstName:         dto.FirstName,
			LastName:          dto.LastName,
			UpdatedAt:         ctime,
			Username:          dto.Username,
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

func (h *UsersHandler) get(w http.ResponseWriter, r *http.Request) {
	accept := r.Header.Get("Accept")
	switch accept {
	case "application/json":
		ret, err := h.rp.GetUsers()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(&ret)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	}
}
