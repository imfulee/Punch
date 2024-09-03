FROM docker.io/golang:1.23-alpine3.20 AS build_stage
WORKDIR /app
RUN apk add just chromium
COPY . .
RUN go mod download
RUN just build

FROM docker.io/alpine:3.20 
RUN apk add chromium
WORKDIR /app 
COPY --from=build_stage /app/punch /app/punch
