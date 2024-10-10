FROM golang:1.21-alpine AS builder

COPY . /github.com/greenblat17/auth/source/
WORKDIR /github.com/greenblat17/auth/source/

RUN go mod download
RUN go build -o ./bin/auth_server cmd/user/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/greenblat17/auth/source/bin/auth_server .
COPY prod.env .

CMD ["./auth_server", "--config-path=prod.env"]