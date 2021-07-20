package main

import (
	"fmt"
)

type MazeGame struct {
}

func (mazeGame MazeGame) createMaze(builder MazeBuilderInterface) (*Maze) {
	builder.buildMaze()

	builder.buildRoom(1)
	builder.buildRoom(2)
	builder.buildRoom(3)
	builder.buildRoom(4)
	builder.buildRoom(5)

	builder.buildDoor(1, 2)
	builder.buildDoor(2, 3)
	builder.buildDoor(3, 4)
	builder.buildDoor(4, 5)

	_, ok := builder.(*StandardMazeBuilder) 
	if (ok){
		r1, ok1 := builder.getMaze().roomNo(1)
		r2, ok2 := builder.getMaze().roomNo(2)

		fmt.Println(r1.getSides(), ok1)
		fmt.Println(r2.getSides(), ok2)
	}

	counterBuilder, ok := builder.(*CountingMazeBuilder)
	if (ok) {
		rooms, doors := counterBuilder.getCounts()
		fmt.Println("Doors:", doors, "Rooms:", rooms)
	}

	return builder.getMaze()
}