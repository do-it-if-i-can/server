package model

import (
	"gorm.io/gorm"
)

// input-----------------------------------------
type NewTodo struct {
	UserID      uint     `json:"user_id"`
	Category    Category `json:"category"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
}

type GetUserByID struct {
	UserID uint `json:"userId"`
}

type GetTodosByCategory struct {
	UserID   uint     `json:"user_id"`
	Category Category `json:"category"`
}

// model-----------------------------------------

type Todo struct {
	gorm.Model
	UserID      uint     `gorm:"not null, column:user_id"`
	Category    Category `gorm:"size:128"`
	Done        bool     `gorm:"default:false"`
	Priority    int64    `gorm:"not null"`
	Title       string   `gorm:"not null"`
	Description *string  ``
	User        User     `gorm:"not null,foreignKey:UserID"`
}
