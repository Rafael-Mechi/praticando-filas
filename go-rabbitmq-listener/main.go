package main

/*
O import eh uma palavra-chave do Go, tipo Class no java.
Coloque quantos imports vc precisar.
O de baixo, temos o alias seguido do link de onde vem esse import.
*/
import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	rabbitmqURL := getEnv("RABBITMQ_URL", "amqp://admin:admin@localhost:5672/")
	queueName := getEnv("QUEUE_NAME", "praticando_fila")
	exchangeName := getEnv("EXCHANGE_NAME", "exchange_praticando_fila")
	routingKey := getEnv("ROUTING_KEY", "praticando_fila.key")

	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		log.Fatalf("Falha ao conectar com RabbitMQ: %s", err)
	}

	defer conn.Close()

	log.Println("Conectado ao RabbitMQ com sucesso!")

	ch, err := conn.Channel()

	if err != nil {
		log.Fatalf("Falha ao abrir canal: %s", err)
	}

	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName, // nome
		"direct",     // tipo
		true,         // durable
		false,        // auto-delete
		false,        // internal
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		log.Fatalf("Falha ao declarar exchange: %s", err)
	}
	log.Printf("Exchange '%s' declarado com sucesso!", exchangeName)

	q, err := ch.QueueDeclare(
		queueName, // nome da fila
		true,      // durable
		false,     // auto-delete
		false,     // exclusive
		false,     // no-wait
		nil,       // args extras
	)

	if err != nil {
		log.Fatalf("Falha ao declarar fila: %s", err)
	}

	log.Printf("Fila '%s' declarada com sucesso!", q.Name)

	err = ch.QueueBind(
		q.Name,       // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Falha ao criar binding: %s", err)
	}
	log.Printf("Binding criado: exchange '%s' -> fila '%s' (key: '%s')", exchangeName, q.Name, routingKey)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}