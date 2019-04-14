package main

import "math/rand"

const blockSize int = 25

// Room represents a room in the game.
type Room struct {
	kind int
}

// NewRoom returns a pointer to a newly allocated Room
func NewRoom() *Room {
	room := new(Room)
	room.kind = rand.Intn(50)

	return room
}
