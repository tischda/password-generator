package main

import (
	"crypto/rand"
)

const dictionary = "abcdefghijklmnopqrstuvxyzABCDEFGHIJKLMNOPQRSTUVXYZ0123456789!@#$%^&*()_+[]{}"

// GeneratePassword generates a password of a specified length.
// The size parameter determines the length of the password.
// It returns the generated password as a string.
func GeneratePassword(size int) (password string) {
	if size > MAX_LENGTH {
		return
	}
	var bytes = make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return err.Error()
	}
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}
