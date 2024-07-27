FROM golang:1.22-alpine as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download & go mod verify

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go
RUN go install github.com/matryer/moq@latest

RUN apk add --no-cache \
    protobuf \
    protobuf-dev \
    protobuf-c \
    protobuf-c-dev

COPY . .

RUN go generate ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o orders ./cmd/orders

FROM scratch
WORKDIR /
COPY --from=build /app/orders /bin/orders
EXPOSE 80
CMD [ "/bin/orders" ]
