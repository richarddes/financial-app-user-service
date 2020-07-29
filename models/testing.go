package models

import (
	"database/sql"
	"log"
)

func UsersExists(db *sql.DB) bool {
	var count int

	queryUser := `SELECT COUNT(id) FROM users;`

	err := db.QueryRow(queryUser).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		return false
	}

	return false
}
