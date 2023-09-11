FROM golang:1.21-alpine

RUN apk add --no-cache tzdata
ENV TZ=Asia/Taipei

WORKDIR /app

COPY . .

RUN echo ' 5 9 * * 1-5 /bin/bash /app/scripts/cron.sh In' >> /etc/crontabs/root
RUN echo ' 55 17 * * 1-5 /bin/bash /app/scripts/cron.sh Out' >> /etc/crontabs/root

RUN go mod download

RUN apk add --no-cache make
RUN make

# Run crond  -f for Foreground 
CMD ["/usr/sbin/crond", "-f"]