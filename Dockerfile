# Etapa de build
FROM golang:1.23.3 AS builder
WORKDIR /app

# Copiar arquivos de dependências e instalar dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar todo o código-fonte do projeto
COPY . .

# Compilar o binário da aplicação
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o temp-service ./cmd/tempservice/main.go

# Etapa de execução

# Etapa final com `alpine` (inclui certificados de CA)
FROM alpine:latest
WORKDIR /app

# Instalar certificados de CA
RUN apk --no-cache add ca-certificates

# Copiar o binário da etapa de build
COPY --from=builder /app/temp-service .

# Copiar arquivos estáticos (exemplo: public/index.html)
COPY --from=builder /app/public ./public

# Copiar variáveis de ambiente
COPY .env .

# Expor a porta usada pelo serviço
EXPOSE 8080

# Entry point do container
ENTRYPOINT ["./temp-service"]
