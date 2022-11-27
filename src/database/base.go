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
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
}

func GenerateConnectionString(c *Credentials) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User,
		c.Password, c.DatabaseName)
}

func Connect(driverName string, c *Credentials) (*sql.DB, error) {
	db, err := sql.Open(driverName, GenerateConnectionString(c))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
