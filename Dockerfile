FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN /go/bin/swag init -g cmd/calculator/main.go

RUN go build -o calculator ./cmd/calculator

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/calculator .
COPY --from=builder /app/docs ./docs

EXPOSE 8080
EXPOSE 8090

CMD ["./calculator"]
