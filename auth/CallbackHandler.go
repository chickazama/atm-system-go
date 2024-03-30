package auth

import (
	"fmt"
	"io"
	"log"
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
	fmt.Println(code)
	url := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", config.ClientID, config.ClientSecret, code)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	fmt.Printf("%s\n", buf)
	url = "/"
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
