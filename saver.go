package main

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Saver struct {
	DB *gorm.DB
}

func (s *Saver) AddHistory(url string) {
	s.DB.Create(&History{
		Url:url,
		CreatedAt:time.Now(),
	})
}

func (s *Saver) GetHistories() []History {
	var histories []History
	s.DB.Find(&histories)
	return histories
}

func (s *Saver) DeleteHistory(id int64) {
	s.DB.Where("id = ?", id).Delete(&History{})
}

