package domain

type User struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Email        string `json:"email" db:"email"`
	Role         string `json:"role" db:"role"`
	Password     string
	PasswordHash []byte `json:"password" db:"password"`
}
