FROM golang:1.22.1 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/myapp .

COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./myapp"]
