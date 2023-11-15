FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

RUN mkdir /app

WORKDIR /app

COPY main.go handlers.go go.mod go.sum ./

RUN go build -o /app/swath ./



# FROM scratch

# COPY --from=builder /app/swath /bin/swath

# CMD ["/bin/swath"]

FROM alpine

RUN apk update add ca-certificates

COPY --from=builder /app/swath /swath

CMD [ "/swath" ]