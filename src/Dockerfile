FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod tidy && go build -o main .

FROM alpine:latest AS runner

WORKDIR /app
COPY --from=builder /app/main .

CMD ["./main"]
