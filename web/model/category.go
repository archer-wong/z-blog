package model

import "time"

type Category struct {
	Id        int    `gorm:"primary_key"`
	Title     string `gorm:"not null;default:'';size:100"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
