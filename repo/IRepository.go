package repo

type IRepository interface {
	UsersTableUp() error
	AccountsTableUp() error
}
