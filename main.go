// Package server starts a mud server.
package main

import (
	"fmt"
	"mud/areas"
	"mud/server"
)

// Start the server.
func main() {

	// Seed the RNG
	//	rand.Seed(time.Now().UnixNano())

	area := areas.NewArea(5, 5)
	fmt.Println("Area", area)

	mapRooms := area.GetMap(15, 15, 5, 5)

	area.DrawAsciiMap(mapRooms)

	database := new(Database)
	database.connect()
	rows, err := database.db.Query("select id from test_table")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		rows.Scan(&id)
		fmt.Println("Id is", id)
	}

	server.StartServer()
}
