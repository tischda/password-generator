package main

import (
	"fmt"
	"strconv"
)

const MAX_LENGTH int = 128

// returns nil if provided length is within boundaries
// or an error if this is not the case.
func lengthValidator(s string) error {
	len, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	return checkLength(len)
}

func checkLength(len int) error {
	if len < 0 || len > MAX_LENGTH {
		return fmt.Errorf("password length must be between 1 and %d", MAX_LENGTH)
	}
	return nil
}
