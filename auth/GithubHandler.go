package auth

import (
	"fmt"
	"net/http"
)

type GithubHandler struct{}

func NewGithubHandler() *GithubHandler {
	return new(GithubHandler)
}

func (h *GithubHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *GithubHandler) get(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("https://github.com/login/oauth/authorize?scope=read:user&client_id=%s&client_secret=%s", config.ClientID, config.ClientSecret)
	http.Redirect(w, r, url, http.StatusSeeOther)
}
