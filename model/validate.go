package model

import (
	"github.com/go-playground/validator"
)

type CustomValidator struct {
	Validator *validator.Validate
}
