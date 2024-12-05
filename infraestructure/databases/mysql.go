package databases

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLConnector struct{}

func (m *MySQLConnector) Connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}
