package ui

import (
	"bufio"
	"os"
)

const (
	OPT_EXIT             = 0
	OPT_MAIN_MENU        = 1
	OPT_SIGNUP_MENU      = 2
	OPT_LOGIN_MENU       = 3
	newlineDelim    byte = '\n'
)

var (
	br = bufio.NewReader(os.Stdin)
)
