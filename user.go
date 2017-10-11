package main

type User struct {
	Username string `gorm:"primary_key;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}
