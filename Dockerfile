FROM golang:alpine AS builder




WORKDIR /app

COPY  go.mod ./
COPY  go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./build/ ./...





#CMD ["/app/build/exchange-rates"]

##
## Deploy
##
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

ENV REDIS_HOST="my_redis"
ENV REDIS_PORT=6379


WORKDIR /app

COPY --from=builder /app/build/* ./

EXPOSE 9000
EXPOSE 9001

CMD ["/app/exchange-rates"]


