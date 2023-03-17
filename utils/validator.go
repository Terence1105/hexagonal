package utils

import (
	"sync"

	"github.com/go-playground/validator"
)

var (
	validate *validator.Validate
	once     sync.Once
)

func NewValidator() {
	once.Do(func() {
		validate = validator.New()
	})
}

func GetValidator() *validator.Validate {
	return validate
}
