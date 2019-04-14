package main

// RoomBlock are blocks of rooms
type RoomBlock struct {
	area  *Area
	rooms [BlockSize][BlockSize]*Room
}

// NewRoomBlock returns a newly allocated RoomBlock
func NewRoomBlock(area *Area) *RoomBlock {
	block := new(RoomBlock)
	block.area = area

	for x := range block.rooms {
		for y := range block.rooms[x] {
			block.rooms[x][y] = NewRoom()
		}
	}

	return block
}

// GetRoom get a specific room in the block
func (rb RoomBlock) GetRoom(x, y int) *Room {
	return rb.rooms[x][y]
}
