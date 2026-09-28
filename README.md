# Praticando Filas com RabbitMQ

## Objetivo

Este projeto tem como objetivo demonstrar o funcionamento de filas de mensagens utilizando RabbitMQ. Ele implementa um cenário simples de producer/consumer onde um serviço em Python envia mensagens para uma fila e um serviço em Go consome essas mensagens.

## Tecnologias

- **RabbitMQ**: Broker de mensagens para gerenciamento de filas
- **Go**: Linguagem utilizada no listener/consumer
- **Python**: Linguagem utilizada no producer
- **Docker**: Containerização do RabbitMQ
- **Docker Compose**: Orquestração do container RabbitMQ

## Requisitos

- Docker e Docker Compose instalados
- Go 1.23+ (para rodar o listener)
- Python 3+ (para rodar o producer)
- Biblioteca `pika` do Python (instalar com `pip install pika`)

## Estrutura de Pastas

```
praticando_fila/
├── compose.yml                          # Configuração do Docker Compose para RabbitMQ
├── go-rabbitmq-listener/               # Consumer em Go
│   ├── main.go                          # Código do listener
│   ├── go.mod                           # Dependências Go
│   └── go.sum                           # Checksum das dependências
├── python-rabitmq-producer/             # Producer em Python
│   └── main.py                          # Código do producer
└── README.md                            # Documentação do projeto
```

## Como Rodar o Projeto

### 1. Iniciar o RabbitMQ

Suba o container do RabbitMQ com Docker Compose:

```bash
docker compose up -d
```

Verifique se o RabbitMQ está rodando:

```bash
docker compose ps
```

O painel administrativo estará disponível em: http://localhost:15672
- Usuário: `admin`
- Senha: `admin`

**Importante:** Antes de rodar os serviços, você precisará criar os seguintes recursos no RabbitMQ através do painel administrativo:
- Uma fila com o nome: `praticando_fila`
- Um exchange com o nome: `exchange_praticando_fila`
- Uma routing key com o nome: `praticando_fila.key`

Certifique-se de configurar o binding entre o exchange e a fila usando a routing key especificada.

### 2. Rodar o Listener (Go)

Navegue até a pasta do listener e execute:

```bash
cd go-rabbitmq-listener
go run main.go
```

O listener ficará aguardando mensagens na fila `praticando_fila`.

### 3. Rodar o Producer (Python)

Em outro terminal, navegue até a pasta do producer e execute:

```bash
cd python-rabitmq-producer
python main.py
```

O producer enviará uma mensagem "Hello World!" para a fila.

### 4. Parar o RabbitMQ

Para parar e remover os containers:

```bash
docker compose down
```
