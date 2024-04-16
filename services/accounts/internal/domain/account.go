package domain

import (
	"errors"
	"net/url"
	"regexp"
	"strings"
	"unicode"

	"github.com/google/uuid"
)

func ValidatePhoneNumber(value string) error {
	if containsNonDigit(value) {
		return errors.New("phone number should contain only digits")
	}

	if len(value) < 8 || len(value) > 15 {
		return errors.New("phone number should be 8 to 15 characters long")
	}

	return nil
}

func ValidatePassword(value string) error {
	if len(value) < 6 {
		return errors.New("password should be at least 8 characters long")
	}

	hasUpperCase := false
	for _, c := range value {
		if unicode.IsUpper(c) {
			hasUpperCase = true
			break
		}
	}

	if !hasUpperCase {
		return errors.New("password should contain at least one upper case letter")
	}

	hasLowerCase := false
	for _, c := range value {
		if unicode.IsLower(c) {
			hasLowerCase = true
			break
		}
	}

	if !hasLowerCase {
		return errors.New("password should contain at least one lower case letter")
	}

	hasDigits := false
	for _, c := range value {
		if unicode.IsDigit(c) {
			hasDigits = true
			break
		}
	}

	if !hasDigits {
		return errors.New("password should contain at least one digit")
	}

	return nil
}

func ValidateName(value string) error {
	if strings.TrimSpace(value) != value {
		return errors.New("name should be trimmed(have no trailling spaces)")
	}

	if strings.ContainsAny(value, "0123456789") {
		return errors.New("name can contain only alphabetical characters")
	}

	if len(value) < 4 {
		return errors.New("name size should be at least 4 characters")
	}

	return nil
}

type Account struct {
	id       uuid.UUID
	phone    string
	password string
	name     *string  // nullable
	profile  *url.URL // nullable
}

func NewAccount(phone string, password string) (*Account, error) {
	account := &Account{id: uuid.New()}

	if err := account.SetPhone(phone); err != nil {
		return nil, err
	}

	if err := account.SetPassword(password); err != nil {
		return nil, err
	}

	return account, nil
}

func (a *Account) ID() uuid.UUID {
	return a.id
}

func (a *Account) Phone() string {
	return a.phone
}

func (a *Account) SetPhone(newPhone string) error {
	if err := ValidatePhoneNumber(newPhone); err != nil {
		return err
	}
	a.password = newPhone
	return nil
}

func (a *Account) Password() string {
	return a.password
}

func (a *Account) SetPassword(newPassword string) error {
	if err := ValidatePassword(newPassword); err != nil {
		return err
	}
	a.password = newPassword
	return nil
}

func (a *Account) Name() string {
	return *a.name
}

func (a *Account) HasName() bool {
	return a.name != nil
}

func (a *Account) SetName(value string) error {
	if err := ValidateName(value); err != nil {
		return err
	}
	a.name = &value
	return nil
}

func (a *Account) HasProfilePicture() bool {
	return a.profile != nil
}

func (a *Account) ProfilePicture() *url.URL {
	return a.profile
}

func (a *Account) SetProfilePicture(u *url.URL) {
	a.profile = u
}

func containsNonDigit(value string) bool {
	pattern := `[^0-9]`
	r := regexp.MustCompile(pattern)
	return r.MatchString(value)
}
