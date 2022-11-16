package postgres

import "database/sql"

func GenerateUsersTable(db *sql.DB) {
	createTable(db, "users", []string{
		"id INT PRIMARY KEY",
		"role_id INT REFERENCES roles(id)",
		"discord BIGINT",
	})
}

func GenerateRolesTable(db *sql.DB) {
	createTable(db, "roles", []string{
		"id SERIAL PRIMARY KEY",
		"name VARCHAR(255)",
		"description text",
	})
}

func GeneratePlacesTable(db *sql.DB) {
	createTable(db, "roles", []string{
		"id SERIAL PRIMARY KEY",
		"name VARCHAR(255)",
	})
}

func GenerateTypesTable(db *sql.DB) {
	createTable(db, "roles", []string{
		"id SERIAL PRIMARY KEY",
		"name VARCHAR(255)",
		"description text",
	})
}

func GenerateActionsTable(db *sql.DB) {
	createTable(db, "roles", []string{
		"id SERIAL PRIMARY KEY",
		"name VARCHAR(255)",
		"place_id INT REFERENCES places(id)",
		"type_id INT REFERENCES types(id)",
		"goal text",
		"data timestamptz",
	})
}
