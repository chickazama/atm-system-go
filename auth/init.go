package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"matthewhope/atm-system-go/services/github"
	"os"
)

var (
	config github.OAuthConfig
	sm     *SessionManager
)

func init() {
	buf, err := os.ReadFile("secrets.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = json.Unmarshal(buf, &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%+v\n", config)
	sm = NewSessionManager()
	go sm.Add("example", "mcgee")
}
