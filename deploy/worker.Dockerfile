FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download || true

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o worker ./cmd/worker/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/worker .

EXPOSE 8000

CMD ["./worker"]
