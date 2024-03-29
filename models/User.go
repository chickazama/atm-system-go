package models

type User struct {
	UserID            int
	CreatedAt         int64
	EmailAddress      string
	EncryptedPassword string
	FirstName         string
	LastName          string
	UpdatedAt         int64
	Username          string
}
