FROM rust:1.80-alpine AS builder

WORKDIR /app

FROM rust:latest AS builder

WORKDIR /app

RUN cargo install sqlx-cli

# Download dependencies 
COPY Cargo.toml Cargo.lock ./
RUN mkdir src
RUN echo "fn main() {}" > src/main.rs
RUN cargo build --release
RUN rm -r src

COPY . .
RUN cargo build --release

FROM scratch
WORKDIR /app
COPY --from=builder /app/target/release/orders /app/orders
COPY --from=builder /app/src/storage/migrations /app/migrations
COPY --from=builder /usr/local/cargo/bin/sqlx /app/sqlx
EXPOSE 80
ENTRYPOINT ["/app/sqlx-cli run --source /app/migrations; /app/orders"]
