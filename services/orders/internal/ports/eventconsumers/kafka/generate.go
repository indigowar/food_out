package kafka

//go:generate mkdir -p ./events
//go:generate protoc --go_out=./events --go_opt=paths=source_relative ./events.proto
