package interfaceInfraestucture

import "database/sql"

type DatabaseConnector interface {
	Connect(dsn string) (*sql.DB, error)
}

type DatabaseConnectorWithOperations interface {
	Connect(dsn string) (interface{}, error) // Retorna una conexión genérica
	Insert(collection string, document interface{}) error
}
