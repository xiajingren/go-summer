package utils

import (
	"encoding/base64"

	"github.com/google/uuid"
)

func NewBase64RandomString() string {
	return base64.StdEncoding.EncodeToString([]byte(uuid.NewString()))
}

func NewUUIdString() string {
	return uuid.NewString()
}
