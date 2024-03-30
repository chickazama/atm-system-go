package models

type Account struct {
	AccountID     int
	AccountNumber string
	Balance       int64
	CreatedAt     int64
	UpdatedAt     int64
	UserID        int
}
