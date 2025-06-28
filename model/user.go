package model

import (
	"time"
)

type User struct {
	Id         int       `json:"id"`
	Email      string    `json:"email" validate:"required, email"`
	Password   string    `json:"password,omitempty" validate:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Deleted_at time.Time `json:"deleted_at"`
}
