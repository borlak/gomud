package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Database connects to a MySQL database
type Database struct {
	db *sql.DB
}

func (database *Database) connect() {
	var err error
	database.db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1)/gomud")
	if err != nil {
		panic(err.Error())
	}

	// Open() doesn't actually open a connection. so we need to validate connection
	err = database.db.Ping()
	if err != nil {
		panic(err.Error())
	}
}
