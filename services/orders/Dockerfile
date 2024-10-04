FROM golang:1.23-alpine AS build

WORKDIR /app

RUN apk add --no-cache \
    protobuf \
    protobuf-dev \
    protobuf-c \
    protobuf-c-dev;

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest;

COPY go.mod go.sum ./

RUN go mod download;

COPY . .

RUN go generate ./...;

RUN CGO_ENABLED=0 GOOS=linux go build -o orders .;


FROM scratch

WORKDIR /

COPY --from=build /app/orders /bin/orders

EXPOSE 80

CMD [ "/bin/orders" ]
