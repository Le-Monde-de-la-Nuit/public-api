package postgres

import (
	"database/sql"
	"strings"
)

func Insert(db *sql.DB, table string, data []any) {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("INSERT INTO " + table + " VALUES " + generateIndexAsString(data))
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(data...)
	if err != nil {
		panic(err)
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}

func createTable(db *sql.DB, name string, columns []string) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + name + " (" + strings.Join(columns, ", ") + ")")
	if err != nil {
		panic(err)
	}
}

func generateIndexAsArray(data []any) []string {
	values := make([]string, len(data))
	for i := 0; i < len(data); i++ {
		values[i] = "$" + string(rune(i+1))
	}
	return values
}

func generateIndexAsString(data []any) string {
	values := generateIndexAsArray(data)
	return strings.Join(values, ", ")
}
