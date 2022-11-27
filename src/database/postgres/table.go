package postgres

func (pg *Postgres) GenerateUsersTable() *Postgres {
	pg.createTable("users", []string{
		"id INT PRIMARY KEY",
		"role_id INT REFERENCES roles(id)",
		"discord BIGINT",
	})
	return pg
}

func (pg *Postgres) GenerateRolesTable() *Postgres {
	pg.createTable("roles", []string{
		"id SERIAL PRIMARY KEY",
		"name VARCHAR(255)",
		"description TEXT",
	})
	return pg
}

func (pg *Postgres) GeneratePlacesTable() *Postgres {
	pg.createTable("places", []string{
		"id SERIAL PRIMARY KEY",
		"name VARCHAR(255)",
	})
	return pg
}

func (pg *Postgres) GenerateTypesTable() *Postgres {
	pg.createTable("types", []string{
		"id SERIAL PRIMARY KEY",
		"name VARCHAR(255)",
		"description TEXT",
	})
	return pg
}

func (pg *Postgres) GenerateActionsTable() *Postgres {
	pg.createTable("actions", []string{
		"id SERIAL PRIMARY KEY",
		"name VARCHAR(255)",
		"place_id INT REFERENCES places(id) ",
		"type_id INT REFERENCES types(id) ",
		"goal TEXT",
		"time TIMESTAMPTZ",
	})
	return pg
}

func (pg *Postgres) GeneratePostsTable() *Postgres {
	pg.createTable("posts", []string{
		"id SERIAL PRIMARY KEY NOT NULL",
		"title VARCHAR(255) NOT NULL",
		"content TEXT NOT NULL",
		"created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL",
	})
	return pg
}

func (pg *Postgres) GenerateTagsTable() *Postgres {
	pg.createTable("tags", []string{
		"id SERIAL PRIMARY KEY NOT NULL",
		"name VARCHAR(255) NOT NULL",
	})
	return pg
}

func (pg *Postgres) GeneratePostTagsTable() *Postgres {
	pg.createTable("post_tags", []string{
		"id SERIAL PRIMARY KEY NOT NULL",
		"post_id INT REFERENCES posts(id) NOT NULL ",
		"tag_id INT REFERENCES tags(id) NOT NULL ",
	})
	return pg
}
