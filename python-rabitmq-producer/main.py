import pika

credentials = pika.PlainCredentials('admin', 'admin')
connection = pika.BlockingConnection(pika.ConnectionParameters(
    host='localhost',
    port=5672,
    credentials=credentials
    ))
channel = connection.channel()

channel.basic_publish(exchange='exchange_praticando_fila',
    routing_key='praticando_fila.key',
    body='Hello World!')


print("Mensagem enviada com sucesso!")