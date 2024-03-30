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
)

type AccountsHandler struct {
	rp repo.IRepository
}

func NewAccountsHandler(r repo.IRepository) *AccountsHandler {
	ret := new(AccountsHandler)
	ret.rp = r
	return ret
}

func (h *AccountsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// CREATE
	case http.MethodPost:
		h.post(w, r)
	}
}

func (h *AccountsHandler) post(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	switch contentType {
	case "application/json":
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		var dto transport.CreateAccountDTO
		err = json.Unmarshal(buf, &dto)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		ctime := time.Now().UTC().UnixMilli()
		a := models.Account{
			AccountNumber: dto.AccountNumber,
			Balance:       dto.Balance,
			CreatedAt:     ctime,
			UpdatedAt:     ctime,
			UserID:        1,
		}
		ret, err := h.rp.CreateAccount(a)
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
