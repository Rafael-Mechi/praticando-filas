package main

/*
O import eh uma palavra-chave do Go, tipo Class.
Coloque quantos imports vc precisar.
O de baixo, temos o alias seguido do link de onde vem esse import.
*/
import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
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

	q, err := ch.QueueDeclare(
		"praticando_fila", // nome da fila
		true,              // durable
		false,             // auto-delete
		false,             // exclusive
		false,             // no-wait
		nil,               // args extras
	)

	if err != nil {
		log.Fatalf("Falha ao declarar fila: %s", err)
	}

	log.Printf("Fila '%s' declarada com sucesso!", q.Name)

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