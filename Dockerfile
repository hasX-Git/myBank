FROM golang:1.21 AS builder

WORKDIR /bankAPI

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bank .

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /bankAPI/bank .

EXPOSE 2266

CMD ["./bank"]