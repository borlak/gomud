package areas

import (
	"errors"
	"fmt"
)

// Area contains roomblocks which contains rooms
type Area struct {
	blocks     [][]*RoomBlock
	maxBlocksX int
	maxBlocksY int
}

// NewArea creates a newly allocated area
func NewArea(blocksX, blocksY int) *Area {
	area := new(Area)
	area.maxBlocksX = blocksX
	area.maxBlocksY = blocksY

	area.blocks = make([][]*RoomBlock, blocksY)
	for i := 0; i < blocksY; i++ {
		area.blocks[i] = make([]*RoomBlock, blocksX)
		for j := 0; j < blocksX; j++ {
			area.blocks[i][j] = NewRoomBlock(area)
		}
	}

	return area
}

// GetBlock gets a specific block using room coordinates
func (area Area) GetBlock(roomX, roomY int) (*RoomBlock, error) {
	if roomX < 0 || roomY < 0 {
		return nil, errors.New("GetBlock roomX or roomY is less than zero")
	}

	blockX := (int)(roomX / BlockSize)
	blockY := (int)(roomY / BlockSize)

	if blockX > area.maxBlocksX || blockY > area.maxBlocksY {
		return nil, errors.New("GetBlock calculated block is outside range")
	}

	return area.blocks[blockX][blockY], nil
}

// GetRoom finds a room in the area given x,y coordinates
func (area Area) GetRoom(roomX, roomY int) (*Room, error) {
	block, err := area.GetBlock(roomX, roomY)
	blockX := (int)(roomX / BlockSize)
	blockY := (int)(roomY / BlockSize)

	if err != nil {
		panic(err)
	}

	room := block.GetRoom(uint16(roomX-(BlockSize*blockX)), uint16(roomY-(BlockSize*blockY)))
	return room, nil
}

// GetMap returns a 2d slice of rooms with their types
func (area Area) GetMap(centerX, centerY, radiusX, radiusY int) [][]uint8 {
	mapRooms := make([][]uint8, radiusX*2+1)
	for i := range mapRooms {
		mapRooms[i] = make([]uint8, radiusY*2+1)
	}

	startX := centerX - radiusX
	startY := centerY - radiusY
	var roomKind uint8

	for x := 0; x < radiusX*2+1; x++ {
		for y := 0; y < radiusY*2+1; y++ {
			roomKind = 0
			currentX := startX + x
			currentY := startY + y

			if currentX < 0 || currentY < 0 {
				roomKind = 0
			} else {
				room, err := area.GetRoom(currentX, currentY)
				if err == nil {
					roomKind = room.kind
				}
			}
			mapRooms[x][y] = roomKind
		}
	}

	return mapRooms
}

func (area Area) DrawAsciiMap(mapRooms [][]uint8) {
	for y := range mapRooms {
		for x := range mapRooms[y] {
			fmt.Printf("[%2d]", mapRooms[x][y])
		}
		fmt.Println()
	}
}
