FROM golang:1.23-alpine AS build

WORKDIR /app

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

RUN apk add --no-cache \
    protobuf \
    protobuf-dev \
    protobuf-c \
    protobuf-c-dev

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go generate ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o order_history .

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/order_history /bin/order_history

EXPOSE 80

CMD [ "/bin/order_history" ]
