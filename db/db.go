package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic(err)
	}

	DB = db
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("cloud not create users table.")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Failed to create events table: " + err.Error())
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		panic("Failed to create registrations table: " + err.Error())
	}
}
