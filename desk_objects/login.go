package objects

import (
	"fmt"
	"regexp"
)

type Login struct {
	Email    string
	Username string
	Password string
}

func ValidateEmail(email string) bool {
	emailFormat, err := regexp.Compile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if err != nil {
		fmt.Println("Error Occurred While Validating Email")
		return false
	}
	if !emailFormat.MatchString(email) {
		fmt.Println("Email is not Formatted correctly")
		return false
	}
	return true
}

func ValidatePassword(password string) bool {
	passwordFormat, err := regexp.Compile(`^(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])(?=.*[@#$%^&+=]).*$`)
	if err != nil {
		fmt.Println("Error Occurred While Validating Password")
		return false
	}
	symbolCheck, err := regexp.Compile(`$&']`)
	if err != nil {
		fmt.Println("Error Occurred While Validating Password")
		return false
	}
	if !symbolCheck.MatchString(password) {
		fmt.Println("Password shouldn't contain certain symbols")
		return false
	}
	if !passwordFormat.MatchString(password) {
		fmt.Println("Invalid Password")
		return false
	}
	return true
}
