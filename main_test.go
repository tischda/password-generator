package main

import (
	"fmt"
	"testing"
)

func TestLengthValidatorValid(t *testing.T) {
	actual := lengthValidator(`19`)

	if nil != actual {
		t.Errorf("Expected: nil, was: %v", actual.Error())
	}
}

func TestLengthValidatorInvalid(t *testing.T) {
	actual := lengthValidator(`-1`)
	expected := fmt.Sprintf("password length must be between 1 and %d", MAX_LENGTH)

	if actual == nil || expected != actual.Error() {
		t.Errorf("Expected: %s, was: %v", expected, actual)
	}
}
