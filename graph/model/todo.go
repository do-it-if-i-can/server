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
	Category    Category `gorm:"not null"`
	Done        bool     `json:"done"`
	Priority    int64    `json:"priority"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	UserID      uint     `json:"user_id"`
	User        User     `json:"user"`
}
