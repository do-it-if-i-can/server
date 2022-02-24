package model

import (
	"time"
)

type NewTodo struct {
	UserID      uint32   `json:"user_id"`
	Category    Category `json:"category"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
}

type GetTodosByCategory struct {
	UserID   uint32   `json:"user_id"`
	Category Category `json:"category"`
}

type Todo struct {
	ID          uint32    `json:"id"`
	Category    Category  `json:"category"`
	Done        bool      `json:"done"`
	Priority    int64     `json:"priority"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      uint32    `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	User        User      `json:"user"`
}
