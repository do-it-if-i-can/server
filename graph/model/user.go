package model

import "time"

type User struct {
	ID          string    `gorm:"primarykey"`
	DisplayName string    `gorm:"not null"`
	UserName    string    `gorm:"not null"`
	Avatar      *string   ``
	Todos       []Todo    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
