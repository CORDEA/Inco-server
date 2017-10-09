package main

import (
	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

const DbName = "inco.db?charset=utf8&parseTime=True&loc=Local"

func main() {
	e := echo.New()

	db, err := gorm.Open("sqlite3", DbName)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&History{})
	if !db.HasTable(&History{}) {
		db.CreateTable(&History{})
	}
	s := Saver{DB:db}

	key, err := ReadPublicKey("")
	if err != nil {
		e.Logger.Fatal(err)
	}
	handler := Handler{
		Key:key,
		Saver:&s,
	}

	e.POST("/histories", handler.PostHistory)
	e.GET("/histories", handler.GetHistories)
	e.Logger.Fatal(e.Start(":8080"))
}