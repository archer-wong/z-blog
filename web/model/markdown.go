package model

import "time"

type Markdown struct {
	Id        int    `gorm:"primary_key"`
	Title     string `gorm:"not null;default:'';size:100"`
	Path      string `gorm:"not null;default:''"`
	ArticleId int    `gorm:"not null;default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
