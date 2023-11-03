package functions

import (
	"errors"
	"regexp"
)

func UsernameValidator(username string) error {
	maxLength := 24

	if len(username) > maxLength {
		return errors.New("username too long")
	}

	return nil

}

func PasswordValidator(password string) error {
	minLength := 8
	maxLength := 24
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString
	hasLowercase := regexp.MustCompile(`[a-z]`).MatchString
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString
	hasSpecialChar := regexp.MustCompile(`[!@#$%^&*()]`).MatchString

	if len(password) < minLength {
		return errors.New("password is too short")
	}
	if len(password) > maxLength {
		return errors.New("password is too long")
	}
	if !hasUppercase(password) {
		return errors.New("password must have an uppercase letter")
	}
	if !hasLowercase(password) {
		return errors.New("password must have a lowercase letter")
	}
	if !hasDigit(password) {
		return errors.New("password must have a digit")
	}
	if !hasSpecialChar(password) {
		return errors.New("password must have a special character")
	}

	return nil
}

func EmailValidator(email string) error {
	maxLength := 24

	if len(email) > maxLength {
		return errors.New("username too long")
	}

	return nil
}
