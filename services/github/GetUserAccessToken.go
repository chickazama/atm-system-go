package github

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetUserAccessToken(code string, config OAuthConfig) {
	url := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", config.ClientID, config.ClientSecret, code)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	req.Header.Set("Accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
		return
	}
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Printf("%s\n", buf)
}
