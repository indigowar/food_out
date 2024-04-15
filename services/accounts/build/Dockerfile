FROM golang:1.21.9 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o accounts ./cmd/accounts


FROM scratch

WORKDIR /

COPY --from=build /app/accounts /bin/accounts

EXPOSE 80

CMD [ "/bin/accounts" ]
