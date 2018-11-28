package model

import "time"

type Link struct {
	Id      int     `gorm:"primary_key"`
	Title   string   `gorm:"not null;default:'';size:100"`
	Url     string   `gorm:"not null;default:''"`
	CreatedAt time.Time
	UpdatedAt time.Time
}


