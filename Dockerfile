# Builder
FROM golang:1.24-alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o backend-exercise ./cmd/main.go

CMD ["/app/backend-exercise"]