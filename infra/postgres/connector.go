package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

const (
	postgresDriverName = "postgres"
	postgresErrorMsg   = "Error while acessing database: "
)

type Connector interface {
	getConnection() (*sqlx.DB, error)
	closeConnection(conn *sqlx.DB)
}

var _ Connector = (*DatabaseConnectionManager)(nil)

type DatabaseConnectionManager struct{}

func (r DatabaseConnectionManager) getConnection() (*sqlx.DB, error) {
	uri, err := GetPostgresConnectionUri()
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Open(postgresDriverName, uri)
	if err != nil {
		log.Print(postgresErrorMsg + err.Error())
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

func (r DatabaseConnectionManager) closeConnection(conn *sqlx.DB) {
	err := conn.Close()
	if err != nil {
		log.Print(err)
	}
}
