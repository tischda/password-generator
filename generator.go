package main

import (
	// TODO: use crypto/rand instead

	"math/rand"
	"time"
)

var characters []string = []string{
	"abcdefghijklmnopqrstuvxyz",
	"ABCDEFGHIJKLMNOPQRSTUVXYZ",
	"0123456789",
	"!@#$%^&*()_+[]{}",
}

func GeneratePassword(passwordLength int) (password string) {
	if passwordLength > MAX_LENGTH {
		return
	}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < passwordLength; i++ {

		// select pool randomly
		n := rand.Intn(len(characters))

		// add random character from pool
		password += getRandom(characters[n])
	}
	return
}

func getRandom(pool string) string {
	i := rand.Intn(len(pool))
	return string(pool[i])
}
