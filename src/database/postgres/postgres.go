package postgres

import (
	"database/sql"
	"database/sql/driver"
	"strings"
	"time"
)

type Postgres struct {
	Db *sql.DB
}

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

func (pg *Postgres) Insert(q string, values ...any) sql.Result {
	db := pg.Db
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	stmt, err := tx.Prepare(q)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(values...)
	if err != nil {
		panic(err)
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
	return result
}

func (pg *Postgres) Get(q string, values ...any) (*sql.Rows, error) {
	db := pg.Db
	result, err := db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (pg *Postgres) createTable(name string, columns []string) {
	_, err := pg.Db.Exec("CREATE TABLE IF NOT EXISTS " + name + " (" + strings.Join(columns, ", ") + ")")
	if err != nil {
		panic(err)
	}
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt *NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}
