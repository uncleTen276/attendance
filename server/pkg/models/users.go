package models

import (
	"time"
)

type User struct {
	ID         int       `db:"id" json:"id"`
	CreatedAt  time.Time `db:"created_at" json:"-"`
	UpdatedAt  time.Time `db:"updated_at" json:"-"`
	Identifier string    `db:"identifier" json:"identifier"`
	Password   string    `db:"password" json:"-"`
	FirstName  string    `db:"first_name" json:"firstName"`
	LastName   string    `db:"last_name" json:"lastName"`
}

func NewUser() *User {
	return &User{}
}

type CreateUser struct {
	Identifier string `json:"identifier" validate:"required,lte=100"`
	Password   string `json:"password" validate:"required,lte=50,gte=8"`
	FirstName  string `json:"firstName" validate:"required,lte=50"`
	LastName   string `json:"lastName" validate:"required,lte=50"`
}

type UserLogin struct {
	Identifier string `json:"identifier" validate:"required,lte=100"`
	Password   string `json:"password" validate:"required,lte=50,gte=8"`
}

type UserToken struct {
	AccessToken string `json:"access_Token"`
}
