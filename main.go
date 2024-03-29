package main

import (
	"log"
	"matthewhope/atm-system-go/ui"
)

func main() {
	var err error
	opt := ui.OPT_MAIN_MENU
	for opt != ui.OPT_EXIT {
		switch opt {
		case ui.OPT_MAIN_MENU:
			opt, err = ui.MainMenu()
			if err != nil {
				log.Fatal(err.Error())
			}
		default:
			log.Fatal("option not recognised")
		}
	}
}
