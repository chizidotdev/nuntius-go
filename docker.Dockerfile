FROM golang:1.21.5-alpine3.19

WORKDIR /app/

RUN apk add --no-cache make
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/a-h/templ/cmd/templ@latest

COPY . /app/.

RUN go mod tidy
