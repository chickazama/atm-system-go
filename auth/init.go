package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var (
	config OAuthConfig
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
}
