package config

import (
	"fmt"
	"os"
)

func GetConnectionStringRabbitMq() (string, error) {
	user, err := getEnv("RABBITMQ_USER")
	if err != nil {
		return "", err
	}

	pass, err := getEnv("RABBITMQ_PASS")
	if err != nil {
		return "", err
	}

	host, err := getEnv("RABBITMQ_HOST")
	if err != nil {
		return "", err
	}

	port, err := getEnv("RABBITMQ_PORT")
	if err != nil {
		return "", err
	}

	/*nameQueueEva, err := getEnv("RABBITMQ_QUEUE_NAME_FILES")
	if err != nil {
		return "", err
	}*/

	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%s", user, pass, host, port)

	return connectionString, nil
}

func getEnv(key string) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}
	return "", fmt.Errorf("variable de entorno no encontrada: %s", key)
}
