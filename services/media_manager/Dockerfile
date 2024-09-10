FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o media_manager .

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/media_manager /bin/media_manager

EXPOSE 80

CMD [ "/bin/media_manager" ]
