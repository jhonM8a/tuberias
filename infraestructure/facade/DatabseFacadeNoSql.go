package facadeDatabase

import (
	"fmt"
	interfaceInfraestucture "tuberias/infraestructure/interface"
)

type DatabaseFacadeNoSql struct {
	connector interfaceInfraestucture.DatabaseConnectorWithOperations
}

func NewDatabaseFacadeNoSql(connector interfaceInfraestucture.DatabaseConnectorWithOperations, dsn string) (*DatabaseFacadeNoSql, error) {
	_, err := connector.Connect(dsn)
	if err != nil {
		return nil, err
	}
	return &DatabaseFacadeNoSql{connector: connector}, nil
}

func (d *DatabaseFacadeNoSql) Insert(collection string, document interface{}) error {
	if d.connector == nil {
		return fmt.Errorf("conector no inicializado")
	}
	return d.connector.Insert(collection, document)
}
