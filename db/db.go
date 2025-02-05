package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("There was some error connecting to the database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {
	query := `CREATE TABLE IF NOT EXISTS tickets (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    registration_number TEXT NOT NULL,
    day1 TEXT NOT NULL,
    day2 TEXT NOT NULL,
    day3 TEST NOT NULL
    )`
	_, err := DB.Exec(query)
	if err != nil {
		panic("There was some error in creating the database")
	}

}
