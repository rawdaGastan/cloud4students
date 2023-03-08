// Package validator for validations
package validator

import (
	"net/mail"

	"github.com/caitlin615/nist-password-validator/password"
)

// ValidateMail used for validating syntax mails
func ValidateMail(address string) error {
	_, err := mail.ParseAddress(address)
	return err
}

// ValidatePassword used for validating passwords before creating user
func ValidatePassword(Password string) error {
	// password should be ASCII , min 5 , max 10
	validator := password.NewValidator(true, 5, 10)
	err := validator.ValidatePassword(Password)
	if err != nil {
		return err
	}
	return nil
}
