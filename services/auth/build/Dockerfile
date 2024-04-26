FROM golang:1.22.2 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o auth ./cmd/auth


FROM scratch

WORKDIR /

COPY --from=build /app/auth /bin/auth

EXPOSE 80

CMD [ "/bin/auth" ]

