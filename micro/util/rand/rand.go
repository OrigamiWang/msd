package main

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

func GenerateRandomString(length int) string {
	buf := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return ""
	}
	keyHex := hex.EncodeToString(buf)
	return keyHex
}
