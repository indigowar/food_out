package service

//go:generate go run github.com/matryer/moq -out storage_moq_test.go . Storage

// TODO: Define the Storage interface.

type Storage interface{}
