package interfaceInfraestucture

import "database/sql"

type DatabaseConnector interface {
	Connect(dsn string) (*sql.DB, error)
}
