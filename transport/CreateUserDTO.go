package transport

type CreateUserDTO struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Username     string `json:"username"`
}
