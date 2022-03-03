package model

import (
	"time"
)

// input-----------------------------------------
type NewTodo struct {
	UserID      string   `json:"user_id"`
	Title       string   `json:"title"`
	Category    Category `json:"category"`
	Description string   `json:"description"`
}

type GetUserByID struct {
	UserID string `json:"userId"`
}

type GetTodosByCategory struct {
	UserID   string   `json:"user_id"`
	Category Category `json:"category"`
}

// model-----------------------------------------

type Todo struct {
	ID          uint      `gorm:"primarykey"`
	UserID      string    `gorm:"not null, column:user_id"`
	Category    Category  `gorm:"size:128"`
	Done        bool      `gorm:"default:false"`
	Priority    int64     `gorm:"not null"`
	Title       string    `gorm:"not null"`
	Description *string   ``
	User        User      `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
