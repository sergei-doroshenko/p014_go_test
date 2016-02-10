package utils

import (
	"crypto/rand"
	"fmt"
)

func GeneratId() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
