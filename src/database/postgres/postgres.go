package postgres

import (
	"database/sql"
	"strings"
)

type Postgres struct {
	Db *sql.DB
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

func createTable(db *sql.DB, name string, columns []string) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + name + " (" + strings.Join(columns, ", ") + ")")
	if err != nil {
		panic(err)
	}
}
