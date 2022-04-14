package main

import (
	"crypto/rand"
)

const dictionary = "abcdefghijklmnopqrstuvxyzABCDEFGHIJKLMNOPQRSTUVXYZ0123456789!@#$%^&*()_+[]{}"

func GeneratePassword(size int) (password string) {
	if size > MAX_LENGTH {
		return
	}
	var bytes = make([]byte, size)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}
