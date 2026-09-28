import pika
import os

# Configurações via variáveis de ambiente
rabbitmq_host = os.getenv('RABBITMQ_HOST', 'localhost')
rabbitmq_port = int(os.getenv('RABBITMQ_PORT', 5672))
rabbitmq_user = os.getenv('RABBITMQ_USER', 'admin')
rabbitmq_pass = os.getenv('RABBITMQ_PASS', 'admin')
queue_name = os.getenv('QUEUE_NAME', 'praticando_fila')
exchange_name = os.getenv('EXCHANGE_NAME', 'exchange_praticando_fila')
routing_key = os.getenv('ROUTING_KEY', 'praticando_fila.key')

credentials = pika.PlainCredentials(rabbitmq_user, rabbitmq_pass)
connection = pika.BlockingConnection(pika.ConnectionParameters(
    host=rabbitmq_host,
    port=rabbitmq_port,
    credentials=credentials
))
channel = connection.channel()

# Declara o exchange
channel.exchange_declare(
    exchange=exchange_name,
    exchange_type='direct',
    durable=True
)
print(f"Exchange '{exchange_name}' declarado com sucesso!")

channel.queue_declare(
    queue=queue_name,
    durable=True
)
print(f"Fila '{queue_name}' declarada com sucesso!")

channel.queue_bind(
    queue=queue_name,
    exchange=exchange_name,
    routing_key=routing_key
)
print(f"Binding criado: exchange '{exchange_name}' -> fila '{queue_name}' (key: '{routing_key}')")

# Publica a mensagem
channel.basic_publish(
    exchange=exchange_name,
    routing_key=routing_key,
    body='Hello World!'
)

print("Mensagem enviada com sucesso!")
connection.close()