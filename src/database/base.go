package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
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

func RowToInterface(r *sql.Row, i interface{}) (interface{}, error) {
	err := r.Scan(i)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func RowsToInterface(r *sql.Rows) (map[string][]interface{}, error) {
	defer r.Close()
	columns, err := r.Columns()
	if err != nil {
		return nil, err
	}
	count := len(columns)
	values := make([]string, count)
	scanArgs := make([]interface{}, count)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	masterData := make(map[string][]interface{})

	for r.Next() {
		err := r.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		for i, v := range values {

			x := v

			//NOTE: FROM THE GO BLOG: JSON and GO - 25 Jan 2011:
			// The json package uses map[string]interface{} and []interface{} values to store arbitrary JSON objects and arrays; it will happily unmarshal any valid JSON blob into a plain interface{} value. The default concrete Go types are:
			//
			// bool for JSON booleans,
			// float64 for JSON numbers,
			// string for JSON strings, and
			// nil for JSON null.

			if nx, ok := strconv.ParseFloat(string(x), 64); ok == nil {
				masterData[columns[i]] = append(masterData[columns[i]], nx)
			} else if b, ok := strconv.ParseBool(string(x)); ok == nil {
				masterData[columns[i]] = append(masterData[columns[i]], b)
			} else if "string" == fmt.Sprintf("%T", string(x)) {
				masterData[columns[i]] = append(masterData[columns[i]], string(x))
			} else {
				fmt.Printf("Failed on if for type %T of %v\n", x, x)
			}

		}
	}
	return masterData, nil
}
