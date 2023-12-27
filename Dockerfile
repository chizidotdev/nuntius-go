FROM golang:1.21.5-alpine3.19

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o main ./cmd/main.go
CMD ["/app/main"]