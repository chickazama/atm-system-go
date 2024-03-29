package ui

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func ProfileMenu() (int, error) {
	br.Reset(os.Stdin)
	err := printTitle()
	if err != nil {
		log.Println(err.Error())
		return OPT_EXIT, err
	}
	fmt.Println("Welcome! Please enter a menu selection:")
	fmt.Println("[1] - Exit")
	buf, err := br.ReadByte()
	if err != nil {
		log.Println(err.Error())
		return OPT_EXIT, err
	}
	opt, err := strconv.Atoi(string(buf))
	if err != nil {
		log.Println(err.Error())
		return OPT_EXIT, err
	}
	switch opt {
	case 1:
		return OPT_EXIT, nil
	default:
		fmt.Println("invalid selection")
		return OPT_MAIN_MENU, nil
	}
}
