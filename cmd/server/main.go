package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")

	// 1. Connection string for the local RabbitMQ Docker container
	connString := "amqp://guest:guest@localhost:5672/"

	// 2. Dial the server to create a connection
	conn, err := amqp.Dial(connString)
	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ: %v", err)
	}
	
	// 3. Defer closing the connection to clean up resources on exit
	defer conn.Close()
	fmt.Println("Connection to RabbitMQ was successful!")

	// 4. Create a channel to listen for system interrupt signals (Ctrl+C)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	// 5. Block the program here until a signal is received
	fmt.Println("Server is running. Press Ctrl+C to shut down.")
	<-signalChan

	// 6. Final shutdown message
	fmt.Println("Shutting down gracefully...")
}