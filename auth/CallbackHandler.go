package auth

import (
	"matthewhope/atm-system-go/services/github"
	"net/http"
)

type CallbackHandler struct {
}

func NewCallbackHandler() *CallbackHandler {
	return new(CallbackHandler)
}

func (h *CallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *CallbackHandler) get(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	github.GetUserAccessToken(code, config)
	url := "/"
	cookie := http.Cookie{
		Name:     "Session",
		Value:    "TestValue",
		Path:     "/",
		MaxAge:   300,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, url, http.StatusSeeOther)
}
