package models

import "time"

type Song struct {
	ID          uint      `gorm:"primaryKey"`
	Group       string    `gorm:"not null" json:"group"`
	Song        string    `gorm:"not null" json:"song"`
	ReleaseDate time.Time `json:"release_date"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
