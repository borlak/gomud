package main

import (
	"mud/areas"
	"strconv"
)

// Mobile is anything living that can move around in the world
type Mobile struct {
	area *areas.Area
	x    uint16
	y    uint16
}

func CMDMove(mob *Mobile, args []string) {
	x, err := strconv.Atoi(args[0])

	if err != nil {
		return
	}
	y, err := strconv.Atoi(args[1])

	if err != nil {
		return
	}

	mob.x = uint16(x)
	mob.y = uint16(y)
}
