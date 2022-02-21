package model

import "time"

type NewTodo struct {
	UserID      string   `json:"user_id"`
	Category    Category `json:"category"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
}

type Todo struct {
	ID          string    `json:"id"`
	Category    Category  `json:"category"`
	Done        bool      `json:"done"`
	Priority    int64     `json:"priority"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      string    `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	User        User      `json:"user"`
}
