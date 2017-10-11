package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

const HmacSecret string = "HmacSecret"

func Generate(code string) string {
	mac := hmac.New(sha256.New, []byte(HmacSecret))
	mac.Write([]byte(code))
	return hex.EncodeToString(mac.Sum(nil))
}
