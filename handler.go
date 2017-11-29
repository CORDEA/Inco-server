package main

import (
	"crypto/rsa"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

const Secret string = "secret"

type Handler struct {
	Key   *rsa.PublicKey
	Saver *Saver
}

func (h *Handler) PostHistory(c echo.Context) error {
	url := c.FormValue("url")
	encrypted, err := Encrypt(h.Key, url)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	encoded := base64.StdEncoding.EncodeToString(encrypted)
	h.Saver.AddHistory(encoded)
	return nil
}

func (h *Handler) DeleteHistory(c echo.Context) error {
	id, err := strconv.ParseInt(c.FormValue("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	h.Saver.DeleteHistory(id)
	return nil
}

func (h *Handler) GetHistories(c echo.Context) error {
	return c.JSON(http.StatusOK, h.Saver.GetHistories())
}

func (h *Handler) Login(c echo.Context) error {
	username := c.QueryParam("user")
	password := c.QueryParam("pass")
	generated := Generate(password)

	user := h.Saver.GetUser(username)
	if generated != user.Password {
		return echo.ErrUnauthorized
	}

	token, err := token(username)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func (h *Handler) SignUp(c echo.Context) error {
	username := c.FormValue("user")
	password := c.FormValue("pass")
	h.Saver.AddUser(username, password)

	token, err := token(username)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func token(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = username
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString([]byte(Secret))
}
