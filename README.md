# Praticando Filas com RabbitMQ

Rafael Mechi de Oliveira - 04251039

## Objetivo

Este projeto tem como objetivo demonstrar o funcionamento de filas de mensagens utilizando RabbitMQ. Ele implementa um cenário simples de producer/consumer onde um serviço em Python envia mensagens para uma fila e um serviço em Go consome essas mensagens.

## Tecnologias

- **RabbitMQ**: Broker de mensagens para gerenciamento de filas
- **Go**: Linguagem utilizada no listener/consumer
- **Python**: Linguagem utilizada no producer
- **Docker**: Containerização de todos os serviços
- **Docker Compose**: Orquestração dos containers

## Requisitos

- Docker e Docker Compose instalados

## Estrutura de Pastas

```
praticando_fila/
├── compose.yml                          # Configuração do Docker Compose para todos os serviços
├── go-rabbitmq-listener/               # Consumer em Go
│   ├── main.go                          # Código do listener
│   ├── go.mod                           # Dependências Go
│   ├── go.sum                           # Checksum das dependências
│   └── Dockerfile                       # Dockerfile do listener
├── python-rabitmq-producer/             # Producer em Python
│   ├── main.py                          # Código do producer
│   ├── requirements.txt                 # Dependências Python
│   └── Dockerfile                       # Dockerfile do producer
└── README.md                            # Documentação do projeto
```

## Como Rodar o Projeto

### 1. Iniciar o RabbitMQ e o Listener

Suba o RabbitMQ e o listener (que ficará aguardando mensagens continuamente):

```bash
docker compose up -d rabbitmq listener
```

Verifique se os serviços estão rodando:

```bash
docker compose ps
```

### 2. Enviar Mensagens

Para enviar mensagens para a fila, rode o producer:

```bash
docker compose run --rm producer
```

### 3. Acompanhar as Mensagens Recebidas

Para ver as mensagens sendo recebidas pelo listener em tempo real:

```bash
docker compose logs -f listener
```

### 4. Parar os Serviços

Para parar e remover os containers:

```bash
docker compose down
```
