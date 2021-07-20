package main

import (
	"fmt"
)

type MazeFactory interface {
	makeMaze() *Maze
	makeRoom(int) Spacer
	makeWall() *Wall
	makeDoor(Spacer, Spacer) Blocker
}

type MazeGame struct {
}

func (mazeGame MazeGame) createMaze() (*Maze) {
	maze := mazeGame.makeMaze()

	r1 := mazeGame.makeRoom(1);
	r2 := mazeGame.makeRoom(2);
	door := mazeGame.makeDoor(r1, r2)

	maze.addRoom(r1)
	maze.addRoom(r2)

	r1.setSide(north, mazeGame.makeWall())
	r1.setSide(east, door)
	r1.setSide(south, mazeGame.makeWall())
	r1.setSide(west, mazeGame.makeWall())

	r2.setSide(north, mazeGame.makeWall())
	r2.setSide(east, mazeGame.makeWall())
	r2.setSide(south, mazeGame.makeWall())
	r2.setSide(west, door)

	fmt.Println(r1.getSides())
	fmt.Println(r2.getSides())
	return maze
}

func (mazeGame MazeGame) makeMaze() *Maze {
	maze := new(Maze)
	maze.rooms = make(map[int]Spacer)

	return maze
}

func (mazeGame MazeGame) makeWall() (*Wall) {
	return new(Wall)
}

func (mazeGame MazeGame) makeRoom(n int) (Spacer){
	room := Room {
		roomNo: n,
	}
	return &room
}

func (mazeGame MazeGame) makeDoor(r1 Spacer, r2 Spacer) (Blocker) {
	room1 := r1.(*Room);
	room2 := r2.(*Room);
	
	door := Door {
		room1: room1,
		room2: room2,
	}
	
	return &door
}


type EnchantedMazeGame struct {
}

func (mazeGame EnchantedMazeGame) createMaze() (*Maze) {
	maze := mazeGame.makeMaze()

	r1 := mazeGame.makeRoom(1);
	r2 := mazeGame.makeRoom(2);
	door := mazeGame.makeDoor(r1, r2)

	maze.addRoom(r1)
	maze.addRoom(r2)

	r1.setSide(north, mazeGame.makeWall())
	r1.setSide(east, door)
	r1.setSide(south, mazeGame.makeWall())
	r1.setSide(west, mazeGame.makeWall())

	r2.setSide(north, mazeGame.makeWall())
	r2.setSide(east, mazeGame.makeWall())
	r2.setSide(south, mazeGame.makeWall())
	r2.setSide(west, door)

	fmt.Println(r1.getSides())
	fmt.Println(r2.getSides())
	enchantedRoom, ok := r1.(*EnchantedRoom)
	if (ok) {
		fmt.Println("Spell:", enchantedRoom.spell)
	}

	return maze
}

func (mazeGame EnchantedMazeGame) makeMaze() (*Maze) {
	maze := new(Maze)
	maze.rooms = make(map[int]Spacer)
	return maze
}

func (mazeGame EnchantedMazeGame) makeWall() (*Wall) {
	return new(Wall)
}

func (mazeGame EnchantedMazeGame) makeRoom(n int) (Spacer){
	room := EnchantedRoom {
		roomNo: n,
		spell: "Alohomora",
	}
	return &room
}

func (mazeGame EnchantedMazeGame) makeDoor(r1 Spacer, r2 Spacer) (Blocker) {
	room1 := r1.(*EnchantedRoom);
	room2 := r2.(*EnchantedRoom);
	
	door := DoorNeedingSpell {
		room1: room1,
		room2: room2,
	}
	return &door
}