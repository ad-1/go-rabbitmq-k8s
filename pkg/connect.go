package pkg

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Connect() (*amqp.Connection, *amqp.Channel) {

	host := GetEnv("RABBITMQ_HOST", "rabbitmq")
	port := GetEnv("RABBITMQ_PORT", "5672")
	user := GetEnv("RABBITMQ_USER", "guest")
	password := GetEnv("RABBITMQ_PASSWORD", "guest")

	url := "amqp://" + user + ":" + password + "@" + host + ":" + port + "/"
	log.Printf("Connecting to RabbitMQ...")
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		log.Fatalf("Failed to open channel: %v", err)
	}

	log.Println("Connected to RabbitMQ")
	return conn, ch
}
