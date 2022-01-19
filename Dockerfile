## Build
FROM golang:alpine AS builder
WORKDIR /app
COPY  go.mod ./
COPY  go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./build/ ./...


## Deploy
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/build/* ./
EXPOSE 9000
EXPOSE 9001



