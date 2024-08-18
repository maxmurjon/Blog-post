package models

import "time"

type Login struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

type User struct{
	Id string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	PasswordHash string `json:"password_hash"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type UserCreateModel struct{
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type UserUpdateModel struct {
    FirstName         *string `json:"first_name,omitempty"`
    LastName          *string `json:"last_name,omitempty"`
    Email             *string `json:"email,omitempty"`
    PasswordHash      *string `json:"password_hash,omitempty"`
}

type PrimaryKey struct{
	Id string `json:"id"`
	Offset int64 `json:"offset,omitempty"`
	Limit int64 `json:"limit,omitempty"`
}