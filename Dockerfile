FROM golang:1.20-alpine AS builder

WORKDIR /app
RUN apk add --no-cache gcc musl-dev sqlite-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

FROM alpine:latest

RUN apk add --no-cache sqlite
WORKDIR /app
COPY --from=builder /app/main .

COPY ./database/bayarind.db /app/database/bayarind.db
EXPOSE 8080

CMD ["./main"]
