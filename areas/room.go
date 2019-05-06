package areas

import "math/rand"

const blockSize int = 25

// Room represents a room in the game.
type Room struct {
	kind uint8
	x    uint16
	y    uint16
}

// NewRoom returns a pointer to a newly allocated Room
func NewRoom(x, y uint16) *Room {
	room := new(Room)
	room.kind = uint8(rand.Intn(50))
	room.x = x
	room.y = y
	return room
}
