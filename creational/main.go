package main

import (
	"errors"
	"fmt"
)

func main() {
	maze := new(Maze)
	maze.rooms = make(map[int]*Room)
	r1 := Room{ 
		roomNo: 0,
	}
	r2 := Room{ 
		roomNo: 1,
	}
	door := Door{
		room1: &r1,
		room2: &r2,
	}

	maze.addRoom(&r1)
	maze.addRoom(&r2)

	r1.setSide(north, new(Wall))
	r1.setSide(east, &door)
	r1.setSide(south, new(Wall))
	r1.setSide(west, new(Wall))

	r2.setSide(north, new(Wall))
	r2.setSide(east, new(Wall))
	r2.setSide(south, new(Wall))
	r2.setSide(west, &door)

	fmt.Println(r1.sides)
	fmt.Println(r2.sides)
}