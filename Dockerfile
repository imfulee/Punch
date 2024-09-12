FROM docker.io/golang:1.23-bookworm AS build_stage
RUN apt-get update -y 
RUN apt-get install -y chromium
RUN apt-get install -y make
WORKDIR /app
COPY . .
RUN go mod download
RUN make

FROM docker.io/debian:bookworm-slim
RUN apt-get update -y 
RUN apt-get install -y chromium
WORKDIR /app 
COPY --from=build_stage /app/punch /app/punch
