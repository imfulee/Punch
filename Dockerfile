FROM docker.io/golang:1.21

WORKDIR /app

COPY . .

RUN apt update -y 
RUN apt install -y make chromium

RUN go mod download

RUN make
