FROM golang:alpine as builder
RUN apk add --no-cache \
    ca-certificates \
    tzdata 
RUN apk add --no-cache git make build-base
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN GO111MODULE=on CGO_ENABLED=1 GOOS=linux go build -mod=vendor -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM alpine:latest 
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /build/main /app/
COPY config.production.yaml /app/
COPY config.development.yaml /app/
#ADD views /app/views/
#ADD certs /app/certs/
#ADD assets /app/assets/
WORKDIR /app
CMD ["./main"]