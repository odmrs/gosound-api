FROM golang:1.22

WORKDIR /usr/src/app

RUN apt-get update && \
    apt-get install -y alsa-tools libasound2 alsa-utils alsa-oss libasound2-dev && \
    rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download 

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main ./cmd/api/

ENV PKG_CONFIG_PATH=/usr/lib/pkgconfig:$PKG_CONFIG_PATH

CMD ["main"]

EXPOSE 4000
