package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type CryptoService struct{}

func (s *CryptoService) HashToken(refreshToken string) (string, error) {
	parts := strings.Split(refreshToken, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid token format")
	}

	payload := parts[1]
	var maxHashedLength int
	if len(payload) < 72 {
		maxHashedLength = len(payload)
	} else {
		maxHashedLength = 72
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(payload[0:maxHashedLength]), 14)

	return string(bytes), err
}
