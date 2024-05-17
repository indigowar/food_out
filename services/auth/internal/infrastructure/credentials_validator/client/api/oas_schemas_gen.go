// Code generated by ogen, DO NOT EDIT.

package api

// Ref: #/components/schemas/AccountCredentials
type AccountCredentials struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// GetPhone returns the value of Phone.
func (s *AccountCredentials) GetPhone() string {
	return s.Phone
}

// GetPassword returns the value of Password.
func (s *AccountCredentials) GetPassword() string {
	return s.Password
}

// SetPhone sets the value of Phone.
func (s *AccountCredentials) SetPhone(val string) {
	s.Phone = val
}

// SetPassword sets the value of Password.
func (s *AccountCredentials) SetPassword(val string) {
	s.Password = val
}

// Ref: #/components/schemas/AccountId
type AccountId struct {
	ID string `json:"id"`
}

// GetID returns the value of ID.
func (s *AccountId) GetID() string {
	return s.ID
}

// SetID sets the value of ID.
func (s *AccountId) SetID(val string) {
	s.ID = val
}

func (*AccountId) validateCredentialsRes() {}

// Ref: #/components/schemas/Error
type Error struct {
	Message string `json:"message"`
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

type ValidateCredentialsBadRequest Error

func (*ValidateCredentialsBadRequest) validateCredentialsRes() {}

type ValidateCredentialsInternalServerError Error

func (*ValidateCredentialsInternalServerError) validateCredentialsRes() {}
