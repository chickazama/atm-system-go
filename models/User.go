package models

type User struct {
	UserID            int    `json:"userID"`
	CreatedAt         int64  `json:"createdAt"`
	EmailAddress      string `json:"emailAddress"`
	EncryptedPassword string `json:"-"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	UpdatedAt         int64  `json:"updatedAt"`
	Username          string `json:"username"`
}
