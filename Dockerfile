FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o app .

FROM debian:buster-slim
WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/assets ./assets
EXPOSE 8080
CMD ["./app"]
