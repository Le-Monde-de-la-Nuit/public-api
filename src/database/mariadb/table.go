package mariadb

func (pg *Mariadb) GenerateUsersTable() *Mariadb {
	pg.createTable("users", []string{
		"id INT PRIMARY KEY",
		"role_id INT REFERENCES roles(id)",
		"discord BIGINT",
	})
	return pg
}

func (pg *Mariadb) GenerateRolesTable() *Mariadb {
	pg.createTable("roles", []string{
		"id INT AUTO_INCREMENT PRIMARY KEY",
		"name VARCHAR(255)",
		"description TEXT",
	})
	return pg
}

func (pg *Mariadb) GeneratePlacesTable() *Mariadb {
	pg.createTable("places", []string{
		"id INT AUTO_INCREMENT PRIMARY KEY",
		"name VARCHAR(255)",
	})
	return pg
}

func (pg *Mariadb) GenerateTypesTable() *Mariadb {
	pg.createTable("types", []string{
		"id INT AUTO_INCREMENT PRIMARY KEY",
		"name VARCHAR(255)",
		"description TEXT",
	})
	return pg
}

func (pg *Mariadb) GenerateActionsTable() *Mariadb {
	pg.createTable("actions", []string{
		"id INT AUTO_INCREMENT PRIMARY KEY",
		"name VARCHAR(255)",
		"place_id INT REFERENCES places(id)",
		"type_id INT REFERENCES types(id)",
		"goal TEXT",
		"time TIMESTAMP",
	})
	return pg
}

func (pg *Mariadb) GeneratePostsTable() *Mariadb {
	pg.createTable("posts", []string{
		"id INT AUTO_INCREMENT PRIMARY KEY ",
		"title VARCHAR(255) ",
		"content LONGTEXT ",
		"created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP()",
	})
	return pg
}

func (pg *Mariadb) GenerateTagsTable() *Mariadb {
	pg.createTable("tags", []string{
		"id INT AUTO_INCREMENT PRIMARY KEY",
		"name VARCHAR(255)",
	})
	return pg
}

func (pg *Mariadb) GeneratePostTagsTable() *Mariadb {
	pg.createTable("post_tags", []string{
		"id INT AUTO_INCREMENT PRIMARY KEY",
		"post_id INT REFERENCES posts(id) ",
		"tag_id INT REFERENCES tags(id) ",
	})
	return pg
}
