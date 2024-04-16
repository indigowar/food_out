package domain

import "testing"

func TestValidatePhoneNumberValidData(t *testing.T) {
	validCases := []string{
		"123456789",
		"79780542142",
		"123125243",
		"12312312312312",
	}

	for _, c := range validCases {
		if err := ValidatePhoneNumber(c); err != nil {
			t.Errorf("domain.ValidatePhoneNumber should not return an error for valid case - %s, but it did: %s", c, err)
		}
	}
}

func TestValidatePhoneNumberInvalidData(t *testing.T) {
	invalidCases := []string{
		"+313 1321 (2312)",
		"1231-322-1321",
		"ababa",
		"123",
		"123213542432123213123123131231312",
		"+ 7(931) 921-32-44",
	}

	for _, c := range invalidCases {
		if ValidatePhoneNumber(c) == nil {
			t.Errorf("domain.ValidatePhoneNumber should return an error, when invalid value: %s is passed", c)
		}
	}
}

func TestValidatePasswordValidData(t *testing.T) {
	validCases := []string{
		"complexPassword123",
		"p234edFCdasdf12",
		"dascvKL123e0zxc",
		"dDd123sda",
	}

	for _, c := range validCases {
		if err := ValidatePassword(c); err != nil {
			t.Errorf("domain.ValidatePassword should not return an error for valid case - %s, but it did: %s", c, err)
		}
	}
}

func TestValidatePasswordInvalidData(t *testing.T) {
	invalidCases := []string{
		"password",
		"not",
		"12345",
		"1234567",
		"AAAAAAH",
		"IbelieveIcanfly",
		"14dollars",
		"14DOLLARS",
	}

	for _, c := range invalidCases {
		if ValidatePassword(c) == nil {
			t.Errorf("domain.ValidatePassword should return an error, when invalid value: %s is passed", c)
		}
	}
}

func TestValidateNameValidData(t *testing.T) {
	validCases := []string{
		"John",
		"Alex",
		"Mark Doe",
		"Johny Peronni",
		"Some guy from Gym",
	}

	for _, c := range validCases {
		if err := ValidateName(c); err != nil {
			t.Errorf("domain.ValidateName should not return an error for valid case - %s, but it did: %s", c, err)
		}
	}
}

func TestValidateNameInvalidData(t *testing.T) {
	invalidCases := []string{
		"d",
		"aaa",
		"	nottrimmed	",
		"je",
		"4131231",
		"   ",
		"",
		"\n\nmax",
		"Johny Peronni 143",
	}

	for _, c := range invalidCases {
		if ValidateName(c) == nil {
			t.Errorf("domain.ValidateName should return an error, when invalid value: %s is passed", c)
		}
	}
}
