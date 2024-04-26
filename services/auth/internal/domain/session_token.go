package domain

import (
	"crypto/rand"
	"encoding/base64"
)

type SessionToken string

func GenerateSessionToken() SessionToken {
	const length = 64

	buffer := make([]byte, length)

	_, err := rand.Read(buffer)
	if err != nil {
		return ""
	}

	return SessionToken(base64.URLEncoding.EncodeToString(buffer)[:length])
}
