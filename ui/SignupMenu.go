package ui

import (
	"fmt"
	"log"
	"matthewhope/atm-system-go/models"
	"matthewhope/atm-system-go/repo"
	"os"
	"time"
)

func SignupMenu(r repo.IRepository) (int, error) {
	err := printTitle()
	if err != nil {
		log.Println(err.Error())
		return OPT_EXIT, err
	}
	u := new(models.User)
	fmt.Println("Please enter the following details:")
	fmt.Print("First Name: ")
	br.Reset(os.Stdin)
	buf, err := br.ReadString(newlineDelim)
	if err != nil {
		log.Println(err.Error())
		return OPT_EXIT, err
	}
	u.FirstName = buf[:len(buf)-1]
	fmt.Print("Last Name: ")
	br.Reset(os.Stdin)
	buf, err = br.ReadString(newlineDelim)
	if err != nil {
		log.Println(err.Error())
		return OPT_EXIT, err
	}
	u.LastName = buf[:len(buf)-1]
	fmt.Print("Email Address: ")
	br.Reset(os.Stdin)
	buf, err = br.ReadString(newlineDelim)
	if err != nil {
		log.Println(err.Error())
		return OPT_EXIT, err
	}
	u.EmailAddress = buf[:len(buf)-1]
	fmt.Print("Username: ")
	br.Reset(os.Stdin)
	buf, err = br.ReadString(newlineDelim)
	if err != nil {
		log.Println(err.Error())
		return OPT_EXIT, err
	}
	u.Username = buf[:len(buf)-1]
	fmt.Print("Password: ")
	br.Reset(os.Stdin)
	buf, err = br.ReadString(newlineDelim)
	if err != nil {
		log.Println(err.Error())
		return OPT_EXIT, err
	}
	u.EncryptedPassword = buf[:len(buf)-1]
	ctime := time.Now().UTC().UnixMilli()
	u.CreatedAt = ctime
	u.UpdatedAt = ctime
	ret, err := r.CreateUser(*u)
	if err != nil {
		return OPT_EXIT, err
	}
	fmt.Printf("%+v\n", ret)
	return OPT_EXIT, nil
}
