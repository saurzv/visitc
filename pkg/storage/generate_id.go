package storage

import (
	"crypto/rand"
	"fmt"
	"strings"
)

// TO-DO: Write better GetId function

func GetId(prefix string) (string, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s_%x", strings.ToLower(prefix), b), nil
}
