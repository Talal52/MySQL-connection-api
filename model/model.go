package model

import "time"

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

type UserUpdate struct {
	Name  string `json:"name,omitempty" db:"name"`
	Email string `json:"email,omitempty" db:"email"`
	Age   int    `json:"age,omitempty" db:"age"`
}
