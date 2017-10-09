package main

import (
	"github.com/labstack/echo"
	"net/http"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
)

type Handler struct {
	Key *rsa.PublicKey
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

func (h *Handler) GetHistories(c echo.Context) error {
	jsonBytes, err := json.Marshal(h.Saver.GetHistories())
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, string(jsonBytes))
}

