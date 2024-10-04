# FROM golang:1.23 AS build
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -o menu .
#
# FROM alpine:latest
# WORKDIR /app
# COPY --from=build /app/menu /bin/menu
# EXPOSE 80
# CMD [ "/bin/menu" ]

FROM golang:1.23-alpine AS build

WORKDIR /app

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest && \
    go install github.com/ogen-go/ogen/cmd/ogen@latest;

COPY go.mod go.sum ./

RUN go mod download;

COPY . .

RUN go generate ./...;

RUN CGO_ENABLED=0 GOOS=linux go build -o menu .;


FROM scratch

WORKDIR /

COPY --from=build /app/menu /bin/menu

EXPOSE 80

CMD [ "/bin/menu" ]
