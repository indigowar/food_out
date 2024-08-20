package commands

import "errors"

var (
	ErrInvalidRequest    = errors.New("incoming request is invalid")
	ErrOrderNotFound     = errors.New("order is not found")
	ErrOrderDuplicated   = errors.New("this order is duplicated, it already exists")
	ErrActionAlreadyDone = errors.New("this action already done to the order")
	ErrUnexpected        = errors.New("unexpected error occurred")
)
