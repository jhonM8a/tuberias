package facadeDatabase

import (
	"database/sql"
	"log"
	interfaceInfraestucture "tuberias/infraestructure/interface"
)

type DatabaseFacade struct {
	DB *sql.DB
}

func NewDatabaseFacade(connector interfaceInfraestucture.DatabaseConnector, dsn string) (*DatabaseFacade, error) {
	db, err := connector.Connect(dsn)
	if err != nil {
		return nil, err
	}
	return &DatabaseFacade{DB: db}, nil
}

func (d *DatabaseFacade) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return d.DB.Query(query, args...)
}

func (d *DatabaseFacade) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.DB.QueryRow(query, args...)
}

func (d *DatabaseFacade) QueryRowByField(query string, fieldValue interface{}, dest ...interface{}) error {
	row := d.DB.QueryRow(query, fieldValue)
	if err := row.Scan(dest...); err != nil {
		return err
	}
	return nil
}

func (d *DatabaseFacade) Close() {
	if err := d.DB.Close(); err != nil {
		log.Printf("Error al cerrar la conexión a la base de datos: %v", err)
	}
}
