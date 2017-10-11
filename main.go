package main

import (
	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/labstack/echo/middleware"
)

const DbName = "inco.db?charset=utf8&parseTime=True&loc=Local"

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := gorm.Open("sqlite3", DbName)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&History{})
	if !db.HasTable(&History{}) {
		db.CreateTable(&History{})
	}
	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}
	s := Saver{DB: db}

	key, err := ReadPublicKey("")
	if err != nil {
		e.Logger.Fatal(err)
	}
	handler := Handler{
		Key:   key,
		Saver: &s,
	}

	e.POST("/signup", handler.SignUp)
	e.GET("/login", handler.Login)

	h := e.Group("/histories")
	h.Use(middleware.JWT([]byte(Secret)))
	h.POST("", handler.PostHistory)
	h.GET("", handler.GetHistories)
	e.Logger.Fatal(e.Start(":8080"))
}
