FROM golang:1.21-alpine

RUN apk add --no-cache tzdata
ENV TZ=Asia/Taipei

WORKDIR /app

COPY . .

RUN cat ./scripts/crontab >> /etc/crontabs/root

RUN go mod download

RUN apk add --no-cache make
RUN make

# Run crond  -f for Foreground 
CMD ["/usr/sbin/crond", "-f"]