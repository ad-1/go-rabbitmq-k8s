package main

import (
	"go-rabbitmq-k8s/pkg"
	"log"
)

func main() {
	numWorkers := pkg.GetEnvAsInt("NUM_WORKERS", 5)
	numMessages := pkg.GetEnvAsInt("NUM_MESSAGES", 5)
	runConsumer(numWorkers, numMessages)
	log.Println("Consumer finished successfully")
}
