package database

import (
	"database/sql"
	"public-api/src/database/mariadb"
)

func GetTagsForPost(db *sql.DB, id int64) ([]mariadb.Tag, error) {
	var (
		tags []mariadb.Tag
		tag  mariadb.Tag
	)
	rows, err := db.Query("SELECT id, name FROM tag WHERE id IN (SELECT tag_id FROM post_tags WHERE post_id = ?)",
		id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return tags, nil
}
