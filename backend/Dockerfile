FROM golang:1.24-alpine

# Adiciona utilitários necessários
RUN apk add --no-cache wget git

# Define diretório de trabalho
WORKDIR /app

# Copia arquivos de dependência primeiro
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

# Expõe a porta do backend (mesma da env APP_PORT)
EXPOSE 3000

# Comando padrão
CMD ["go", "run", "main.go"]