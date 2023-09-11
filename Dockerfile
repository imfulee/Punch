FROM golang:1.21

WORKDIR /app

COPY . .

RUN go mod download
RUN make

RUN apt-get update && apt-get install -y cron
RUN crontab ./scripts/cron
