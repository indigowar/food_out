// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package queries

import (
	"context"
	"github.com/google/uuid"
	"github.com/indigowar/food_out/services/menu/domain"
	"sync"
)

// Ensure, that MenuRetrieverMock does implement MenuRetriever.
// If this is not the case, regenerate this file with moq.
var _ MenuRetriever = &MenuRetrieverMock{}

// MenuRetrieverMock is a mock implementation of MenuRetriever.
//
//	func TestSomethingThatUsesMenuRetriever(t *testing.T) {
//
//		// make and configure a mocked MenuRetriever
//		mockedMenuRetriever := &MenuRetrieverMock{
//			RetrieveByIDFunc: func(ctx context.Context, id uuid.UUID) (*domain.Menu, error) {
//				panic("mock out the RetrieveByID method")
//			},
//			RetrieveByRestaurantFunc: func(ctx context.Context, restaurant uuid.UUID) ([]*domain.Menu, error) {
//				panic("mock out the RetrieveByRestaurant method")
//			},
//		}
//
//		// use mockedMenuRetriever in code that requires MenuRetriever
//		// and then make assertions.
//
//	}
type MenuRetrieverMock struct {
	// RetrieveByIDFunc mocks the RetrieveByID method.
	RetrieveByIDFunc func(ctx context.Context, id uuid.UUID) (*domain.Menu, error)

	// RetrieveByRestaurantFunc mocks the RetrieveByRestaurant method.
	RetrieveByRestaurantFunc func(ctx context.Context, restaurant uuid.UUID) ([]*domain.Menu, error)

	// calls tracks calls to the methods.
	calls struct {
		// RetrieveByID holds details about calls to the RetrieveByID method.
		RetrieveByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
		// RetrieveByRestaurant holds details about calls to the RetrieveByRestaurant method.
		RetrieveByRestaurant []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Restaurant is the restaurant argument value.
			Restaurant uuid.UUID
		}
	}
	lockRetrieveByID         sync.RWMutex
	lockRetrieveByRestaurant sync.RWMutex
}

// RetrieveByID calls RetrieveByIDFunc.
func (mock *MenuRetrieverMock) RetrieveByID(ctx context.Context, id uuid.UUID) (*domain.Menu, error) {
	if mock.RetrieveByIDFunc == nil {
		panic("MenuRetrieverMock.RetrieveByIDFunc: method is nil but MenuRetriever.RetrieveByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockRetrieveByID.Lock()
	mock.calls.RetrieveByID = append(mock.calls.RetrieveByID, callInfo)
	mock.lockRetrieveByID.Unlock()
	return mock.RetrieveByIDFunc(ctx, id)
}

// RetrieveByIDCalls gets all the calls that were made to RetrieveByID.
// Check the length with:
//
//	len(mockedMenuRetriever.RetrieveByIDCalls())
func (mock *MenuRetrieverMock) RetrieveByIDCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockRetrieveByID.RLock()
	calls = mock.calls.RetrieveByID
	mock.lockRetrieveByID.RUnlock()
	return calls
}

// RetrieveByRestaurant calls RetrieveByRestaurantFunc.
func (mock *MenuRetrieverMock) RetrieveByRestaurant(ctx context.Context, restaurant uuid.UUID) ([]*domain.Menu, error) {
	if mock.RetrieveByRestaurantFunc == nil {
		panic("MenuRetrieverMock.RetrieveByRestaurantFunc: method is nil but MenuRetriever.RetrieveByRestaurant was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		Restaurant uuid.UUID
	}{
		Ctx:        ctx,
		Restaurant: restaurant,
	}
	mock.lockRetrieveByRestaurant.Lock()
	mock.calls.RetrieveByRestaurant = append(mock.calls.RetrieveByRestaurant, callInfo)
	mock.lockRetrieveByRestaurant.Unlock()
	return mock.RetrieveByRestaurantFunc(ctx, restaurant)
}

// RetrieveByRestaurantCalls gets all the calls that were made to RetrieveByRestaurant.
// Check the length with:
//
//	len(mockedMenuRetriever.RetrieveByRestaurantCalls())
func (mock *MenuRetrieverMock) RetrieveByRestaurantCalls() []struct {
	Ctx        context.Context
	Restaurant uuid.UUID
} {
	var calls []struct {
		Ctx        context.Context
		Restaurant uuid.UUID
	}
	mock.lockRetrieveByRestaurant.RLock()
	calls = mock.calls.RetrieveByRestaurant
	mock.lockRetrieveByRestaurant.RUnlock()
	return calls
}