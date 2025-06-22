package db

import (
	"database/sql"
	"log"
)

func RunMigrations(db *sql.DB) {
	queries := []string{
		`
			CREATE TABLE IF NOT EXISTS tickets (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				title TEXT NOT NULL,
				description TEXT NOT NULL,
				client_id INTEGER NOT NULL,
				status TEXT NOT NULL,
				created_at DATETIME NOT NULL,
				updated_at DATETIME NOT NULL
			);
		`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(`Error running migration: %v`, err)
		}
	}

	log.Println("Migrations executed successfully")
}
