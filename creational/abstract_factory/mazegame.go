package main

import (
	"fmt"
)

type MazeGame struct {
}

func (mazeGame MazeGame) createMaze(factory MazeFactoryInterface) (*Maze) {
	maze := factory.makeMaze()
	r1 := factory.makeRoom(1)
	r2 := factory.makeRoom(2)
	door := factory.makeDoor(r1, r2)

	maze.addRoom(r1)
	maze.addRoom(r2)

	r1.setSide(north, factory.makeWall())
	r1.setSide(east, door)
	r1.setSide(south, factory.makeWall())
	r1.setSide(west, factory.makeWall())

	r2.setSide(north, factory.makeWall())
	r2.setSide(east, factory.makeWall())
	r2.setSide(south, factory.makeWall())
	r2.setSide(west, door)

	fmt.Println(r1.getSides())
	fmt.Println(r2.getSides())

	spell, ok := r2.(*EnchantedRoom)
	if (ok){
		fmt.Println(spell.spell)
	}

	return maze
}