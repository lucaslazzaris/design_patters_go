package main

import (
	"fmt"
)

type MazeGame struct {
}

func (mazeGame MazeGame) CreateMaze(factory MazeFactoryInterface) (*Maze) {
	maze := factory.MakeMaze()
	r1 := factory.MakeRoom(1)
	r2 := factory.MakeRoom(2)
	door := factory.MakeDoor(r1, r2)

	maze.addRoom(r1)
	maze.addRoom(r2)

	r1.setSide(north, factory.MakeWall())
	r1.setSide(east, door)
	r1.setSide(south, factory.MakeWall())
	r1.setSide(west, factory.MakeWall())

	r2.setSide(north, factory.MakeWall())
	r2.setSide(east, factory.MakeWall())
	r2.setSide(south, factory.MakeWall())
	r2.setSide(west, door)

	fmt.Println(r1.getSides())
	fmt.Println(r2.getSides())

	spell := r2.(*EnchantedRoom)
	fmt.Println(spell.spell)

	return maze
}