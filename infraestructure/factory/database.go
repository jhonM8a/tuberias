package factory

import (
	"fmt"
	"tuberias/infraestructure/databases"
	interfaceInfraestucture "tuberias/infraestructure/interface"
)

type DatabaseFactory struct{}

const (
	databaseMySQL = "mysql"
)

func (f *DatabaseFactory) GetDatabaseConnector(dbType string) (interfaceInfraestucture.DatabaseConnector, error) {
	switch dbType {
	case databaseMySQL:
		return &databases.MySQLConnector{}, nil
	default:
		return nil, fmt.Errorf("tipo de base de datos no soportado: %s", dbType)
	}
}

func (f *DatabaseFactory) GetDatabaseConnectorNoSQL(dbType string) (interfaceInfraestucture.DatabaseConnectorWithOperations, error) {
	switch dbType {
	case "mongo":
		return &databases.MongoDBConnector{}, nil
	default:
		return nil, fmt.Errorf("tipo de base de datos no soportado: %s", dbType)
	}
}
