package model

import (
	"time"
)

// model-----------------------------------------
type Todo struct {
	ID          uint      `gorm:"primarykey"`
	UserID      string    `gorm:"index:idx_user_id, sort:desc, column:user_id, not null"`
	Category    Category  `gorm:"size:64"`
	Done        bool      `gorm:"default:false"`
	Priority    int64     `gorm:"not null"`
	Title       string    `gorm:"not null"`
	Description *string   ``
	User        User      `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

// input-----------------------------------------
type GetTodosByCategory struct {
	UserID   string   `json:"user_id"`
	Category Category `json:"category"`
}

type GetTodosByUser struct {
	UserID string
}

type NewTodo struct {
	UserID      string   `json:"user_id"`
	Title       string   `json:"title"`
	Category    Category `json:"category"`
	Description string   `json:"description"`
}

type UpdateTodo struct {
	TodoID      uint
	Title       string
	Description string
}
type UpdateTodoDone struct {
	TodoID uint
	Done   bool
}

type DeleteTodo struct {
	TodoID uint
}

type CopyTodo struct {
	TodoID uint
}

type MoveTodo struct {
	TodoID       uint
	Category     Category
	AfterTodoIds []uint
}
