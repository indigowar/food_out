FROM golang:1.23-alpine AS build

WORKDIR /app

RUN go install github.com/ogen-go/ogen/cmd/ogen@latest;

COPY go.mod go.sum ./

RUN go mod download;

COPY . .

RUN go generate ./...;

RUN CGO_ENABLED=0 GOOS=linux go build -o auth ./cmd/auth


FROM scratch

WORKDIR /

COPY --from=build /app/auth /bin/auth

EXPOSE 80

CMD [ "/bin/auth" ]
