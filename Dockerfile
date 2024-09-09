FROM docker.io/golang:1.23-bookworm as build_stage
WORKDIR /app
COPY . .
RUN apt-get update -y 
RUN apt-get install -y make chromium
RUN go mod download
RUN make

FROM docker.io/debian:bookworm 
RUN apt-get update -y 
RUN apt-get install -y chromium
WORKDIR /app 
COPY --from=build_stage /app/punch /app/punch