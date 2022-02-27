package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string  `gorm:"not null"`
	Avatar *string ``
	Todos  []Todo  `gorm:"not null,references:UserID;"`
}
