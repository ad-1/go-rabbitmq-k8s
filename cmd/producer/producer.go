package main

import (
	"fmt"
	"go-rabbitmq-k8s/pkg"
	"log"
	"math/rand"

	amqp "github.com/rabbitmq/amqp091-go"
)

func runProducer(numMessages int) error {
	log.Printf("Running producer to publish %d messages", numMessages)

	conn, ch := pkg.DeclareQueue() // Ensure the queue is declared before publishing messages
	defer conn.Close()
	defer ch.Close()

	for i := range numMessages {
		body := getRandomName()
		err := ch.Publish(
			"",            // default exchange
			"hello-queue", // routing key = queue name
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			},
		)
		if err != nil {
			return err
		}
		log.Printf("Published message %d: %s", i, body)
	}

	log.Printf("Finished publishing %d messages", numMessages)
	return nil
}

var names = []string{
	"emma", "olivia", "ava", "isabella", "sophia",
	"liam", "noah", "elijah", "james", "william",
	"mia", "amelia", "harper", "evelyn", "abigail",
	"benjamin", "lucas", "henry", "alexander", "mason",
}

func getRandomName() string {
	name := names[rand.Intn(len(names))]
	return fmt.Sprintf("name:%s", name)
}
