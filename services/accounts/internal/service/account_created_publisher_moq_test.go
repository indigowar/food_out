// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package service

import (
	"context"
	"github.com/indigowar/food_out/services/accounts/internal/domain"
	"sync"
)

// Ensure, that AccountCreatedPublisherMock does implement AccountCreatedPublisher.
// If this is not the case, regenerate this file with moq.
var _ AccountCreatedPublisher = &AccountCreatedPublisherMock{}

// AccountCreatedPublisherMock is a mock implementation of AccountCreatedPublisher.
//
//	func TestSomethingThatUsesAccountCreatedPublisher(t *testing.T) {
//
//		// make and configure a mocked AccountCreatedPublisher
//		mockedAccountCreatedPublisher := &AccountCreatedPublisherMock{
//			PublishAccountCreatedFunc: func(ctx context.Context, account *domain.Account) error {
//				panic("mock out the PublishAccountCreated method")
//			},
//		}
//
//		// use mockedAccountCreatedPublisher in code that requires AccountCreatedPublisher
//		// and then make assertions.
//
//	}
type AccountCreatedPublisherMock struct {
	// PublishAccountCreatedFunc mocks the PublishAccountCreated method.
	PublishAccountCreatedFunc func(ctx context.Context, account *domain.Account) error

	// calls tracks calls to the methods.
	calls struct {
		// PublishAccountCreated holds details about calls to the PublishAccountCreated method.
		PublishAccountCreated []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Account is the account argument value.
			Account *domain.Account
		}
	}
	lockPublishAccountCreated sync.RWMutex
}

// PublishAccountCreated calls PublishAccountCreatedFunc.
func (mock *AccountCreatedPublisherMock) PublishAccountCreated(ctx context.Context, account *domain.Account) error {
	if mock.PublishAccountCreatedFunc == nil {
		panic("AccountCreatedPublisherMock.PublishAccountCreatedFunc: method is nil but AccountCreatedPublisher.PublishAccountCreated was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Account *domain.Account
	}{
		Ctx:     ctx,
		Account: account,
	}
	mock.lockPublishAccountCreated.Lock()
	mock.calls.PublishAccountCreated = append(mock.calls.PublishAccountCreated, callInfo)
	mock.lockPublishAccountCreated.Unlock()
	return mock.PublishAccountCreatedFunc(ctx, account)
}

// PublishAccountCreatedCalls gets all the calls that were made to PublishAccountCreated.
// Check the length with:
//
//	len(mockedAccountCreatedPublisher.PublishAccountCreatedCalls())
func (mock *AccountCreatedPublisherMock) PublishAccountCreatedCalls() []struct {
	Ctx     context.Context
	Account *domain.Account
} {
	var calls []struct {
		Ctx     context.Context
		Account *domain.Account
	}
	mock.lockPublishAccountCreated.RLock()
	calls = mock.calls.PublishAccountCreated
	mock.lockPublishAccountCreated.RUnlock()
	return calls
}