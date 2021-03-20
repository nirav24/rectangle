#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /go/src/app/rectangle-assessment .

RUN chmod +x /app/rectangle-assessment
LABEL Name=rectangle Version=0.0.1
