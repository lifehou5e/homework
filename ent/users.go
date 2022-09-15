package ent

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lifehou5e/homework/servergorilla/validation"
)

type Users struct {
	ID        uuid.UUID
	Email     string `json:"email"`
	Password  string `json:"password"`
	FullName  string `json:"fullname"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *Users) Validation(user Users) error {
	if validation.IsASCII(u.Password) == false {
		return errors.New("password must contain only ASCII symbols")
	}
	if len(u.Password) < 8 {
		return errors.New("password should have at least 8 characters and be less than 256 symbols")
	}
	if len(u.Password) >= 256 {
		return errors.New("password should have less than 256 characters")
	}
	if len(u.Email) >= 256 {
		return errors.New("email should have less than 256 characters")
	}
	if !validation.ContainDog(u.Email) {
		return errors.New("invalid email")
	}

	return nil
}
