package repo

import "matthewhope/atm-system-go/models"

type IRepository interface {
	UsersTableUp() error
	AccountsTableUp() error
	CreateUser(models.User) (models.User, error)
	CreateAccount(models.Account) (models.Account, error)
	// GetUsers() ([]models.User, error)
}
