package ent

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID        uuid.UUID
	Email     string `json:"email"`
	Password  string `json:"password"`
	FullName  string `json:"fullname"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
