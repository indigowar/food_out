FROM golang:1.22.3 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o menu .

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/menu /bin/menu

EXPOSE 80

CMD [ "/bin/menu" ]
