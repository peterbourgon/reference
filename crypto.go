package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Inspiration from https://github.com/micro/user-srv

const fixed = "0ca2838077a795941b44fa14b2125a97140e51f0"

func salt(plaintextPassword string) (salt, digest string) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	salt = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	hash, err := bcrypt.GenerateFromPassword([]byte(fixed+salt+plaintextPassword), 16)
	if err != nil {
		panic(err)
	}
	digest = base64.StdEncoding.EncodeToString(hash)

	return salt, digest
}
