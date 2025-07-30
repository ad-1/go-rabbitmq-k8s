package main

import (
	"go-rabbitmq-k8s/cmd/consumer/processor"
	"go-rabbitmq-k8s/pkg"
	"log"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

func runConsumer(numWorkers, numMessages int) {

	log.Printf("Running consumer with %d workers, processing %d messages", numWorkers, numMessages)

	conn, ch := pkg.DeclareQueue()
	defer conn.Close()
	defer ch.Close()

	// QoS: only send this many unacked messages at once
	ch.Qos(numWorkers, 0, false)

	msgs, err := ch.Consume(
		"hello-queue",
		"",
		true,  // auto-ack
		false, // not exclusive
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	// Internal Go channel between your message consumer and your worker goroutines
	jobChan := make(chan amqp.Delivery)
	var wg sync.WaitGroup

	processor := processor.NewMockAgifyProcessor() // Create a processor instance

	// Launch workers
	for i := range numWorkers {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, jobChan, processor) // workers will exit when they finish their current job
		}(i)
	}

	log.Println("Waiting for messages to dispatch...")
	dispatchedMessages := 0
	for msg := range msgs {
		jobChan <- msg
		dispatchedMessages++
		if dispatchedMessages >= numMessages {
			log.Printf("Dispatched %d messages, closing job channel", dispatchedMessages)
			close(jobChan) // Close the job channel to signal workers to stop
			break
		}
	}

	wg.Wait() // Wait for all workers to finish
	log.Println("All workers have finished processing")

}

func worker(id int, jobs <-chan amqp.Delivery, processor processor.Processor) {
	for job := range jobs {
		body := string(job.Body)
		log.Printf("Worker %d processing message: %s", id, body)
		result, err := processor.Process(string(job.Body))
		if err != nil {
			log.Printf("Worker %d failed to process '%s': %v", id, body, err)
			continue
		}
		log.Printf("Worker %d processed message '%s' with result: %s", id, body, result)
	}
}
