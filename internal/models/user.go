package models

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint64       `json:"id" db:"id" validate:"omitempty"`
	UserName  string       `json:"username" db:"username" validate:"omitempty"`
	Login     string       `json:"login" db:"login" validate:"required"`
	Password  string       `json:"password" db:"password" validate:"required"`
	CreatedAt time.Time    `json:"created_at" db:"created_at" validate:"omitempty"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at" validate:"omitempty"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePasswords(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

func (u *User) SanitizePassword() {
	u.Password = ""
}

type UserWithToken struct {
	UserName string `json:"username"`
	Token    string `json:"token"`
}
