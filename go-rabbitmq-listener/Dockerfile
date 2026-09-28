# Stage de build
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copia os arquivos de dependências
COPY go.mod go.sum ./

# Baixa as dependências
RUN go mod download

# Copia o código fonte
COPY main.go .

# Compila o binário
RUN CGO_ENABLED=0 GOOS=linux go build -o listener main.go

# Stage final
FROM alpine:latest

WORKDIR /app

# Copia o binário compilado
COPY --from=builder /app/listener .

# Executa o binário
CMD ["./listener"]
