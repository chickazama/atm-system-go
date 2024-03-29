package main

import (
	"log"
	"matthewhope/atm-system-go/repo"
	"matthewhope/atm-system-go/ui"
)

func main() {
	r := repo.NewSQLiteRepository()
	var err error
	opt := ui.OPT_MAIN_MENU
	for opt != ui.OPT_EXIT {
		switch opt {
		case ui.OPT_MAIN_MENU:
			opt, err = ui.MainMenu()
			if err != nil {
				log.Fatal(err.Error())
			}
		case ui.OPT_SIGNUP_MENU:
			opt, err = ui.SignupMenu(r)
			if err != nil {
				log.Fatal(err.Error())
			}
		case ui.OPT_PROFILE_MENU:
			opt, err = ui.ProfileMenu()
			if err != nil {
				log.Fatal(err.Error())
			}
		default:
			log.Fatal("option not recognised")
		}
	}
}
