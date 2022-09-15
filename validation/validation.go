package validation

import (
	"errors"
	"unicode"
)

func IsASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func FullNameValidation(fn string) error {
	if len(fn) < 2 {
		return errors.New("full Name is too short")
	}
	return nil
}

func ContainDog(email string) bool {
	for _, v := range email {
		if string(v) == "@" {
			return true
		}
	}
	return false
}
