package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
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

	e.POST("/histories", handler.PostHistory)

	h := e.Group("/histories")
	h.Use(middleware.JWT([]byte(Secret)))
	h.GET("", handler.GetHistories)
	h.DELETE("", handler.DeleteHistory)

	// FIXME: Deprecated
	// You should use DELETE.
	// Red language has not implemented DELETE yet. So, POST is available.
	d := e.Group("/delete-histories")
	d.Use(middleware.JWT([]byte(Secret)))
	d.POST("", handler.DeleteHistory)

	e.Logger.Fatal(e.Start(":8080"))
}
