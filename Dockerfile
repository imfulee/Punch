FROM docker.io/golang:1.21

WORKDIR /app

COPY . .

RUN apt-get update -y 
RUN apt-get install -y make chromium

RUN go mod download

RUN make
