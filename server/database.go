package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Database connects to a MySQL database
type Database struct {
	connection *sql.DB
}

func (database *Database) connect() {
	var err error
	database.connection, err = sql.Open("mysql", "root:password@tcp(127.0.0.1)/gomud")
	if err != nil {
		panic(err.Error())
	}

	// Open() doesn't open a connection. validate connection
	err = database.connection.Ping()
	if err != nil {
		panic(err.Error())
	}
}
