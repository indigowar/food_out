FROM golang:1.22.5 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o orders .

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/orders /bin/orders

EXPOSE 80

CMD [ "/bin/orders" ]
