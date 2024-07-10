// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/url"

	"github.com/google/uuid"
)

// Ref: #/components/schemas/Dish
type Dish struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Image url.URL   `json:"image"`
	Price float64   `json:"price"`
}

// GetID returns the value of ID.
func (s *Dish) GetID() uuid.UUID {
	return s.ID
}

// GetName returns the value of Name.
func (s *Dish) GetName() string {
	return s.Name
}

// GetImage returns the value of Image.
func (s *Dish) GetImage() url.URL {
	return s.Image
}

// GetPrice returns the value of Price.
func (s *Dish) GetPrice() float64 {
	return s.Price
}

// SetID sets the value of ID.
func (s *Dish) SetID(val uuid.UUID) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *Dish) SetName(val string) {
	s.Name = val
}

// SetImage sets the value of Image.
func (s *Dish) SetImage(val url.URL) {
	s.Image = val
}

// SetPrice sets the value of Price.
func (s *Dish) SetPrice(val float64) {
	s.Price = val
}

func (*Dish) retrieveDishByIDRes() {}

// Ref: #/components/schemas/Error
type Error struct {
	Code    float64 `json:"code"`
	Message string  `json:"message"`
}

// GetCode returns the value of Code.
func (s *Error) GetCode() float64 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *Error) SetCode(val float64) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

func (*Error) retrieveListOfRestaurantsRes() {}

// Ref: #/components/schemas/Menu
type Menu struct {
	ID         uuid.UUID   `json:"id"`
	Name       string      `json:"name"`
	Restaurant uuid.UUID   `json:"restaurant"`
	Image      url.URL     `json:"image"`
	Dishes     []uuid.UUID `json:"dishes"`
}

// GetID returns the value of ID.
func (s *Menu) GetID() uuid.UUID {
	return s.ID
}

// GetName returns the value of Name.
func (s *Menu) GetName() string {
	return s.Name
}

// GetRestaurant returns the value of Restaurant.
func (s *Menu) GetRestaurant() uuid.UUID {
	return s.Restaurant
}

// GetImage returns the value of Image.
func (s *Menu) GetImage() url.URL {
	return s.Image
}

// GetDishes returns the value of Dishes.
func (s *Menu) GetDishes() []uuid.UUID {
	return s.Dishes
}

// SetID sets the value of ID.
func (s *Menu) SetID(val uuid.UUID) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *Menu) SetName(val string) {
	s.Name = val
}

// SetRestaurant sets the value of Restaurant.
func (s *Menu) SetRestaurant(val uuid.UUID) {
	s.Restaurant = val
}

// SetImage sets the value of Image.
func (s *Menu) SetImage(val url.URL) {
	s.Image = val
}

// SetDishes sets the value of Dishes.
func (s *Menu) SetDishes(val []uuid.UUID) {
	s.Dishes = val
}

func (*Menu) retrieveMenuByIDRes() {}

type RetrieveDishByIDInternalServerError Error

func (*RetrieveDishByIDInternalServerError) retrieveDishByIDRes() {}

type RetrieveDishByIDNotFound Error

func (*RetrieveDishByIDNotFound) retrieveDishByIDRes() {}

type RetrieveListOfRestaurantsOKApplicationJSON []uuid.UUID

func (*RetrieveListOfRestaurantsOKApplicationJSON) retrieveListOfRestaurantsRes() {}

type RetrieveMenuByIDInternalServerError Error

func (*RetrieveMenuByIDInternalServerError) retrieveMenuByIDRes() {}

type RetrieveMenuByIDNotFound Error

func (*RetrieveMenuByIDNotFound) retrieveMenuByIDRes() {}

type RetrieveMenusByRestaurantInternalServerError Error

func (*RetrieveMenusByRestaurantInternalServerError) retrieveMenusByRestaurantRes() {}

type RetrieveMenusByRestaurantNotFound Error

func (*RetrieveMenusByRestaurantNotFound) retrieveMenusByRestaurantRes() {}

type RetrieveMenusByRestaurantOKApplicationJSON []Menu

func (*RetrieveMenusByRestaurantOKApplicationJSON) retrieveMenusByRestaurantRes() {}

type ValidateRestaurantDishesBadRequest Error

func (*ValidateRestaurantDishesBadRequest) validateRestaurantDishesRes() {}

type ValidateRestaurantDishesInternalServerError Error

func (*ValidateRestaurantDishesInternalServerError) validateRestaurantDishesRes() {}

type ValidateRestaurantDishesNotFound Error

func (*ValidateRestaurantDishesNotFound) validateRestaurantDishesRes() {}

// ValidateRestaurantDishesOK is response for ValidateRestaurantDishes operation.
type ValidateRestaurantDishesOK struct{}

func (*ValidateRestaurantDishesOK) validateRestaurantDishesRes() {}

// Ref: #/components/schemas/ValidationList
type ValidationList struct {
	Restaurant uuid.UUID   `json:"restaurant"`
	Dishes     []uuid.UUID `json:"dishes"`
}

// GetRestaurant returns the value of Restaurant.
func (s *ValidationList) GetRestaurant() uuid.UUID {
	return s.Restaurant
}

// GetDishes returns the value of Dishes.
func (s *ValidationList) GetDishes() []uuid.UUID {
	return s.Dishes
}

// SetRestaurant sets the value of Restaurant.
func (s *ValidationList) SetRestaurant(val uuid.UUID) {
	s.Restaurant = val
}

// SetDishes sets the value of Dishes.
func (s *ValidationList) SetDishes(val []uuid.UUID) {
	s.Dishes = val
}