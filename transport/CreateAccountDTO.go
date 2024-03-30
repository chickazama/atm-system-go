package transport

type CreateAccountDTO struct {
	AccountNumber string `json:"accountNumber"`
	Balance       int64  `json:"balance"`
}
