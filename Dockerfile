FROM golang:1.24.3

WORKDIR /bankAPI

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bank ./main

EXPOSE 2266

ENTRYPOINT ["./bank"]