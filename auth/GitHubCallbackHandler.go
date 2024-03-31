package auth

import (
	"log"
	"matthewhope/atm-system-go/services/github"
	"net/http"
)

type GitHubCallbackHandler struct {
}

func NewGitHubCallbackHandler() *GitHubCallbackHandler {
	return new(GitHubCallbackHandler)
}

func (h *GitHubCallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *GitHubCallbackHandler) get(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	value, err := github.GetUserAccessToken(code, config)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "unable to retrieve code", http.StatusInternalServerError)
		return
	}
	url := "/"
	cookie := http.Cookie{
		Name:     "Session",
		Value:    value,
		Path:     "/",
		MaxAge:   300,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, url, http.StatusSeeOther)
}
