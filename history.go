package main

import (
	"time"
)

type History struct {
	ID int64 `gorm:"primary_key"`
	Url string `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}
