FROM golang:1.22.2 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o menu ./cmd/menu

FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN apk add --no-cache curl && \
    curl -LO https://github.com/pressly/goose/releases/download/v3.5.3/goose_3.5.3_linux_amd64.tar.gz && \
    tar -C /usr/local/bin -xzf goose_3.5.3_linux_amd64.tar.gz && \
    rm goose_3.5.3_linux_amd64.tar.gz

WORKDIR /app

COPY --from=build /app/internal/ports/storage/postgres/migrations ./migrations
COPY --from=build /app/menu /bin/menu

EXPOSE 80

CMD goose -dir /migrations postgres "user=$POSTGRES_USER password=$POSTGRES_PASSWORD dbname=$POSTGRES_DBNAME host=$POSTGRES_HOST port=$POSTGRES_PORT" up && /bin/menu
