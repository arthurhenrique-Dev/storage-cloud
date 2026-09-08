package models

type User struct {
	ID           int64  `json:"id"`
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	PasswordHash string `json:"-"`
}

func NewUser(email, firstName, lastName, passwordHash string) *User {
	return &User{
		Email:        email,
		FirstName:    firstName,
		LastName:     lastName,
		PasswordHash: passwordHash,
	}
}

