package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")

	if err != nil {
		panic(fmt.Sprint("Could not connect to database because: ", err))
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = createTables()
	if err != nil {
		panic(fmt.Sprint("Could not create Events table, because: ", err))
	}
}

func createTables() error {
	createUsersTableQuery := `
	CREATE TABLE IF NOT EXISTS Users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTableQuery)
	if err != nil {
		panic(fmt.Sprint("Could not create Users table because ", err))
	}

	createEventsTableQuery := `
	CREATE TABLE IF NOT EXISTS Events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES Users(id)
	)
	`

	_, err = DB.Exec(createEventsTableQuery)
	if err != nil {
		panic(fmt.Sprint("Could not create Events table because ", err))
	}

	return nil
}
