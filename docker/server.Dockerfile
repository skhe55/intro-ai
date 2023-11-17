FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go

FROM alpine:latest

RUN apk update && apk upgrade

RUN rm -rf /var/cache/apk/* && \
    rm -rf /tmp/* && \
    apk add --no-cache make

RUN adduser -D ikbkr

WORKDIR /app

COPY --from=builder /app /app
RUN mkdir ftpdata && chown ikbkr -R ./ftpdata/

ENV GOOSE_VERSION_TAG="v3.16.0"
ADD "https://github.com/pressly/goose/releases/download/$GOOSE_VERSION_TAG/goose_linux_x86_64" /bin/goose
RUN chmod +x /bin/goose

USER ikbkr

CMD ["./app"]