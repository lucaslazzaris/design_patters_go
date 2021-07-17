package main

import (
	"fmt"
)

type MazeGame struct {
}

func (mazeGame MazeGame) CreateMaze(builder MazeBuilderInterface) (*Maze) {
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

	fmt.Println(builder.getMaze().rooms)
	

	counterBuilder, ok := builder.(*CountingMazeBuilder)
	if (ok) {
		rooms, doors := counterBuilder.getCounts()
		fmt.Println("Doors:", doors, "Rooms:", rooms)
	}

	return builder.getMaze()
}