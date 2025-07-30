package main

import (
	"go-rabbitmq-k8s/pkg"
	"log"
)

func main() {
	numMessages := pkg.GetEnvAsInt("NUM_MESSAGES", 10)
	if err := runProducer(numMessages); err != nil {
		log.Fatalf("Failed to run producer: %v", err)
	}
	log.Println("Producer finished successfully")
}
