FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

RUN mkdir /app

WORKDIR /app

COPY main.go handlers.go go.mod go.sum ./

RUN go build -o /app/swath ./



FROM alpine

RUN apk update add --no-cache ca-certificates

COPY --from=builder /app/swath  /app/

COPY htmx /app/htmx

WORKDIR /app

ARG version 

ENV SWATH_VERSION $version

CMD [ "./swath" ]