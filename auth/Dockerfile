FROM golang:1.19.1-alpine3.15 AS builder
COPY . /go/src/projects/LDmitryLD/hugoproxy-microservices/auth
WORKDIR /go/src/projects/LDmitryLD/hugoproxy-microservices/auth

RUN go build -ldflags="-w -s" -o /go/bin/server /go/src/projects/LDmitryLD/hugoproxy-microservices/auth/cmd/app

FROM alpine:3.15

COPY --from=builder /go/bin/server /go/bin/server
COPY ./.env /app/.env

WORKDIR /app

ENTRYPOINT [ "/go/bin/server" ]