# Etapa de compilação
FROM golang:1.22-alpine as builder

RUN apk add --no-cache git alsa-lib-dev gcc musl-dev

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download 

COPY . .
RUN go build -v -o main ./cmd/api/

FROM alpine:latest

RUN apk update && \
    apk add --no-cache alsa-lib mplayer

WORKDIR /app

COPY --from=builder /usr/src/app/main .

ENV PKG_CONFIG_PATH=/usr/lib/pkgconfig:$PKG_CONFIG_PATH

CMD ["./main"]

EXPOSE 4000
