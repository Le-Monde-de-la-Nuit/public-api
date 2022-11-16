package postgres

import (
	"database/sql"
	"strings"
)

func Insert(db *sql.DB, q string, values ...any) sql.Result {
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

func createTable(db *sql.DB, name string, columns []string) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + name + " (" + strings.Join(columns, ", ") + ")")
	if err != nil {
		panic(err)
	}
}
