// Package server starts a mud server.
package main

import (
	"fmt"
)

// Start the server.
func main() {

	// Seed the RNG
	//	rand.Seed(time.Now().UnixNano())

	area := NewArea(5, 5)
	fmt.Println("Area", area)

	/*
		room, err := area.GetBlock(25, 25)

		if err != nil {
			panic(err)
		}
		fmt.Println("Checking for room 25,25", room)
	*/

	mapRooms := area.GetMap(15, 15, 5, 5)

	area.DrawAsciiMap(mapRooms)

	StartServer()
}
