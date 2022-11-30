package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	PublicCredentials *Credentials
)

type Credentials struct {
	User     string
	Password string
}

func (c *Credentials) GenerateConnectionString(db string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", c.User,
		c.Password, "db-"+db, 3306, db)
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
