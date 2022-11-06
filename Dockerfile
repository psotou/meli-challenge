# syntax=docker/dockerfile:1

## Build
FROM golang:1.17-buster AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# RUN go build -o /challenge-meli
RUN go build -o /challenge-meli ./cmd/server/

## Deploy
FROM gcr.io/distroless/base-debian10 AS deployer

WORKDIR /

COPY --from=builder /challenge-meli .

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./challenge-meli"]
