package validation

import (
	"errors"
	"unicode"
)

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func PasswordValidation(pass string) error {
	if isASCII(pass) == false {
		return errors.New("password must contain only ASCII symbols")
	}
	if len(pass) < 8 {
		return errors.New("password should have at least 8 characters and be less than 256 symbols")
	}
	if len(pass) >= 256 {
		return errors.New("password should have less than 256 characters")
	}
	return nil
}

func EmailValidation(email string) error {
	if len(email) >= 256 {
		return errors.New("email should have less than 256 characters")
	}

	if !containDog(email) {
		return errors.New("invalid email")
	}

	return nil
}

func FullNameValidation(fn string) error {
	if len(fn) < 2 {
		return errors.New("full Name is too short")
	}
	return nil
}

func containDog(email string) bool {
	for _, v := range email {
		if string(v) == "@" {
			return true
		}
	}
	return false
}
