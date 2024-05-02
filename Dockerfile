# Etapa de compilação
FROM golang:1.22-alpine as builder

# Instale dependências de compilação necessárias
RUN apk add --no-cache git alsa-lib-dev gcc musl-dev

WORKDIR /usr/src/app

# Copie e baixe as dependências do Go primeiro para aproveitar o cache das camadas Docker
COPY go.mod go.sum ./
RUN go mod download 

# Copie o código fonte e compile
COPY . .
RUN go build -v -o main ./cmd/api/

# Etapa final
FROM alpine:latest

# Instale apenas as bibliotecas necessárias no runtime
RUN apk update && \
    apk add --no-cache alsa-lib mplayer

WORKDIR /app

# Copie o executável do builder
COPY --from=builder /usr/src/app/main .

# Configure variáveis de ambiente, se necessário
ENV PKG_CONFIG_PATH=/usr/lib/pkgconfig:$PKG_CONFIG_PATH

# Defina o comando para executar o aplicativo
CMD ["./main"]

# Exponha a porta necessária
EXPOSE 4000
