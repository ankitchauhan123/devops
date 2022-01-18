FROM golang:alpine AS builder


ENV REDIS_HOST="my_redis"
ENV REDIS_PORT=6379


WORKDIR /app

COPY  go.mod ./
COPY  go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./build/ ./...

EXPOSE 9000
EXPOSE 9001



CMD ["/app/build/exchange-rates"]

##
## Deploy
##
#FROM debian:buster-slim

#RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
#    ca-certificates && \
#    rm -rf /var/lib/apt/lists/*

#WORKDIR /app

#COPY --from=builder /app/build/* ./

#EXPOSE 8080

#CMD ["exchange-rates"]


