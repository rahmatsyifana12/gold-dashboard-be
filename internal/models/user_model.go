package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Email    string `json:"email"`
	Password string `json:"-"`
	Name     string `json:"name"`
}
