package main

import (
	"fmt"
	"log"
	"tuberias/config"
	broker "tuberias/infraestructureFacade"
)

func main() {
	fmt.Println(" <---- Inicio de tuberias ---->")

	// [1] Cargar propiedades
	fmt.Println(" <-- [1] Cargar propiedades -->")

	connectionString, err := config.GetConnectionStringRabbitMq()
	if err != nil {
		log.Fatalf("Error al obtener la cadena de conexión: %v", err)
	}
	fmt.Println("RabbitMQ Connection String:", connectionString)

	// [2] Crear conexión
	fmt.Println(" <-- [2] Crear conexión  -->")
	brokerType := "rabbitmq" // Podría cambiar a "kafka"

	b, err := broker.NewBroker(brokerType, connectionString)
	if err != nil {
		log.Fatalf("Error creating broker: %v", err)
	}
	defer b.Close()

	if err := b.Connect(); err != nil {
		log.Fatalf("Error connecting to broker: %v", err)
	}

	// Escuchar mensajes continuamente
	queueName := "test_queue"
	msgChan, err := b.Consume(queueName)
	if err != nil {
		log.Fatalf("Error al consumir mensajes de la cola '%s': %v", queueName, err)
	}

	fmt.Printf("Esperando mensajes en la cola '%s'...\n", queueName)

	// Bucle infinito para procesar los mensajes
	for msg := range msgChan {
		fmt.Printf("Mensaje recibido: %s\n", string(msg))
		// Aquí puedes agregar lógica para procesar el mensaje.
	}
}
