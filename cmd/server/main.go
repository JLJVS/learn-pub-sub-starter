package main

import "fmt"
import (
    amqp "github.com/rabbitmq/amqp091-go"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("Starting Peril server...")

	// Connection string for RabbitMQ
    connectionString := "amqp://guest:guest@localhost:5672/"

	// step 2 start amqp dial
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		// Handle error - perhaps log it and exit the program
		fmt.Printf("Failed to connect to RabbitMQ: %v", err)
		os.Exit(1)
	}

	// step 3 defer
	// Defer closing the connection
	defer conn.Close()

	// step 4 Print connection success message
	fmt.Println("Successfully connected to RabbitMQ")

	// step 5
	// Wait for ctrl+c
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	// step 6
	// If a signal is received, print shutdown message
	fmt.Println("Interrupt received, shutting down...")
	// The connection will be closed automatically by the deferred conn.Close()

}
