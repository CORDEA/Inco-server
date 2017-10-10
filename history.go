package main

import (
	"time"
)

type History struct {
	ID int64 `gorm:"primary_key" json:"id"`
	Url string `gorm:"not null" json:"url"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}
