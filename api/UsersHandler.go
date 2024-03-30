package api

import "matthewhope/atm-system-go/repo"

type UsersHandler struct {
	rp repo.IRepository
}

func NewUsersHandler(r repo.IRepository) *UsersHandler {
	ret := new(UsersHandler)
	ret.rp = r
	return ret
}
