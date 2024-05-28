FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o pos-go-cleanarch-listorders cmd/main.go

FROM postgres:13-alpine as db
ENV POSTGRES_DB=order_db
ENV POSTGRES_USER=user
ENV POSTGRES_PASSWORD=secret
COPY db/migrations /docker-entrypoint-initdb.d/

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/pos-go-cleanarch-listorders .
COPY --from=db /docker-entrypoint-initdb.d /docker-entrypoint-initdb.d
EXPOSE 8080
CMD ["./pos-go-cleanarch-listorders"]
