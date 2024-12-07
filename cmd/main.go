package main

import (
	"fmt"
	"log"
	config "tuberias/config"
	broker "tuberias/infraestructure/factory"
	filterMetadata "tuberias/services/impl"
)

func main() {
	fmt.Println(" <---- Inicio de tuberias ---->")

	// [1] Cargar propiedades
	fmt.Println(" <-- [1] Cargar propiedades -->")

	config, err := config.GetConnectionStringRabbitMq()
	if err != nil {
		log.Fatalf("Error al obtener la cadena de conexión: %v", err)
	}
	fmt.Println("RabbitMQ Connection String:", config.ConnectionString)

	// [2] Crear conexión
	fmt.Println(" <-- [2] Crear conexión  -->")

	b, err := broker.NewBroker(config.BrokerName, config.ConnectionString)
	if err != nil {
		log.Fatalf("Error creating broker: %v", err)
	}
	defer b.Close()

	if err := b.Connect(); err != nil {
		log.Fatalf("Error connecting to broker: %v", err)
	}

	// Escuchar mensajes continuamente
	msgChan, err := b.Consume(config.QueueName)
	if err != nil {
		log.Fatalf("Error al consumir mensajes de la cola '%s': %v", config.QueueName, err)
	}

	fmt.Printf("Esperando mensajes en la cola '%s'...\n", config.QueueName)

	// Bucle infinito para procesar los mensajes
	for msg := range msgChan {
		fmt.Printf("Mensaje recibido: %s\n", string(msg))
		// Aquí puedes agregar lógica para procesar el mensaje.
		filterMetadata.FiletMetadata(string(msg))
	}
}
