package model

import "time"

type User struct {
	ID        uint      `gorm:"primarykey"`
	Name      string    `gorm:"not null"`
	Avatar    *string   ``
	Todos     []Todo    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
