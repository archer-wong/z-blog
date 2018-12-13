package model

import "time"

type Article struct {
	Id         int    `gorm:"primary_key"`
	Title      string `gorm:"not null;default:'';size:100"`
	Content    string `gorm:"type:text;not null"`
	Tag        string `gorm:"default:'';size:100"`
	View       int    `gorm:"default:0"`
	CategoryId int    `gorm:"not null;default:0"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Category   Category
}

type ArticleWithCategory struct {
	Article      Article
	CategoryName string
}
