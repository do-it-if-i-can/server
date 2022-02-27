package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string  `json:"name"`
	Avatar *string `json:"avatar"`
	Todos  []Todo  `json:"todos"`
}
