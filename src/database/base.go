package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var (
	PublicCredentials *Credentials
)

type Credentials struct {
	User     string
	Password string
}

func (c *Credentials) GenerateConnectionString(db string) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "db-"+db, 5432, c.User,
		c.Password, db)
}

func (c *Credentials) Connect(driverName string, dbName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, c.GenerateConnectionString(dbName))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
