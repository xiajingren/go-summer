package utils

import (
	"encoding/base64"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func NewBase64RandomString() string {
	return base64.StdEncoding.EncodeToString([]byte(uuid.NewString()))
}

func NewUUIdString() string {
	return uuid.NewString()
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Fatalf("an error occurred while hash and salt %w \n", err)
	}
	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}
