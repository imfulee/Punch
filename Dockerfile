FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN apk add --no-cache chromium make

RUN make