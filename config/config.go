package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

type ConfigConnectionBroker struct {
	ConnectionString string
	QueueName        string
	BrokerName       string
}

type ConfigConnectionDatabase struct {
	ConnectionString string
	Database         string
}

type ConfigConnectionDatabaseNoSQL struct {
	ConnectionString string
	Database         string
}

func GetConnectionStringRabbitMq() (ConfigConnectionBroker, error) {
	user, err := getEnv("RABBITMQ_USER")
	if err != nil {
		return ConfigConnectionBroker{}, err
	}

	pass, err := getEnv("RABBITMQ_PASS")
	if err != nil {
		return ConfigConnectionBroker{}, err
	}

	host, err := getEnv("RABBITMQ_HOST")
	if err != nil {
		return ConfigConnectionBroker{}, err
	}

	port, err := getEnv("RABBITMQ_PORT")
	if err != nil {
		return ConfigConnectionBroker{}, err
	}

	nameQueueEva, err := getEnv("RABBITMQ_QUEUE_NAME_FILES")
	if err != nil {
		return ConfigConnectionBroker{}, err
	}

	brokerName, err := getEnv("BROKER_NAME")

	if err != nil {
		return ConfigConnectionBroker{}, err
	}

	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%s", user, pass, host, port)

	configConnection := ConfigConnectionBroker{
		ConnectionString: connectionString,
		QueueName:        nameQueueEva,
		BrokerName:       brokerName,
	}

	return configConnection, nil
}

func GetConnectionDatabse() (ConfigConnectionDatabase, error) {
	user, err := getEnv("DB_USER")
	if err != nil {
		return ConfigConnectionDatabase{}, err
	}

	pass, err := getEnv("DB_PASSWORD")
	if err != nil {
		return ConfigConnectionDatabase{}, err
	}

	host, err := getEnv("DB_HOST")
	if err != nil {
		return ConfigConnectionDatabase{}, err
	}

	port, err := getEnv("DB_PORT")
	if err != nil {
		return ConfigConnectionDatabase{}, err
	}

	database, err := getEnv("DB_DATABASENAME")
	if err != nil {
		return ConfigConnectionDatabase{}, err
	}

	databaseEngine, err := getEnv("DB_ENGINE")
	if err != nil {
		return ConfigConnectionDatabase{}, err
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, database)

	configConnectionDatabase := ConfigConnectionDatabase{
		ConnectionString: connectionString,
		Database:         databaseEngine,
	}

	return configConnectionDatabase, nil

}

func GetConnectionDatabaseNoSQL() (ConfigConnectionDatabaseNoSQL, error) {
	// Obtén las variables de entorno específicas de MongoDB
	user, err := getEnv("NOSQL_DB_USER") // root
	if err != nil {
		return ConfigConnectionDatabaseNoSQL{}, err
	}

	pass, err := getEnv("NOSQL_DB_PASSWORD") // rootpassword123
	if err != nil {
		return ConfigConnectionDatabaseNoSQL{}, err
	}

	host, err := getEnv("NOSQL_DB_HOST") // 181.79.9.72
	if err != nil {
		return ConfigConnectionDatabaseNoSQL{}, err
	}

	port, err := getEnv("NOSQL_DB_PORT") // 6451
	if err != nil {
		return ConfigConnectionDatabaseNoSQL{}, err
	}

	database, err := getEnv("NOSQL_NAME") // admin
	if err != nil {
		return ConfigConnectionDatabaseNoSQL{}, err
	}

	// Construir la cadena de conexión de MongoDB
	connectionString := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/?authMechanism=SCRAM-SHA-1&authSource=admin",
		user, pass, host, port,
	)

	configConnectionDatabaseNoSQL := ConfigConnectionDatabaseNoSQL{
		ConnectionString: connectionString,
		Database:         database,
	}

	return configConnectionDatabaseNoSQL, nil
}

func getEnv(key string) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}
	return "", fmt.Errorf("variable de entorno no encontrada: %s", key)
}

// Leer y cargar la clave privada desde un archivo PEM
func readPEMFile(fileName string) ([]byte, error) {
	// Leer el archivo
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error al leer el archivo %s: %v", fileName, err)
	}
	return data, nil
}

// Función para cargar una clave privada RSA desde un archivo PEM
func LoadPrivateKeyFromFile(fileName string) (*rsa.PrivateKey, error) {
	// Leer el contenido del archivo PEM
	data, err := readPEMFile(fileName)
	if err != nil {
		return nil, err
	}

	// Decodificar el contenido PEM
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("error decodificando el archivo PEM: %s", fileName)
	}

	// Parsear la clave privada RSA
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error al parsear la clave privada: %v", err)
	}

	return privateKey, nil
}

// Función para cargar una clave pública RSA desde un archivo PEM
func LoadPublicKeyFromFile(fileName string) (*rsa.PublicKey, error) {
	// Leer el contenido del archivo PEM
	data, err := readPEMFile(fileName)
	if err != nil {
		return nil, err
	}

	// Decodificar el contenido PEM
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("error decodificando el archivo PEM: %s", fileName)
	}

	// Parsear la clave pública RSA
	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error al parsear la clave pública: %v", err)
	}

	return publicKey, nil
}
