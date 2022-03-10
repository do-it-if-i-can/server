package model

import "time"

// model-----------------------------------------
type User struct {
	ID          string    `gorm:"primarykey"`
	DisplayName string    `gorm:"not null"`
	UserName    string    `gorm:"not null"`
	Avatar      *string   ``
	Todos       []Todo    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

// input-----------------------------------------

type GetUserByID struct {
	UserID string
}

type UpsertUser struct {
	UserID      string
	DisplayName string
	UserName    string
	Avatar      string
}
