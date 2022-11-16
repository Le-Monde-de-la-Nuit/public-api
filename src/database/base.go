package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
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

func ParseConnectionString(c string) *Credentials {
	splited := strings.Split(c, " ")
	credentials := &Credentials{}
	for _, v := range splited {
		values := strings.Split(v, "=")
		switch values[0] {
		case "host":
			credentials.Host = values[1]
		case "port":
			i, err := strconv.Atoi(values[1])
			if err != nil {
				panic(err)
			}
			credentials.Port = i
		case "user":
			credentials.User = values[1]
		case "password":
			credentials.Password = values[1]
		case "dbname":
			credentials.DatabaseName = values[1]
		}
	}
	return credentials
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

func RowsToInterface(r *sql.Rows, i []interface{}) ([]interface{}, error) {
	defer r.Close()
	err := r.Scan(i...)
	if err != nil {
		return nil, err
	}
	return i, nil
}
