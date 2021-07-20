package main

type MazeFactoryInterface interface {
	makeMaze() *Maze
	makeWall() *Wall
	makeRoom(int) Spacer 
	makeDoor(Spacer, Spacer) Blocker
}

type MazeFactory struct {
}

func (mazeFactory *MazeFactory) makeMaze() (*Maze) {
	var maze Maze
	maze.rooms = make(map[int]Spacer)
	return &maze
}

func (mazeFactory *MazeFactory) makeWall() (*Wall) {
	return new(Wall)
}

func (mazeFactory *MazeFactory) makeRoom(n int) (Spacer){
	room := Room {
		roomNo: n,
	}
	return &room
}

func (mazeFactory *MazeFactory) makeDoor(r1 Spacer, r2 Spacer) (Blocker) {
	room1 := r1.(*Room);
	room2 := r2.(*Room);
	
	door := Door {
		room1: room1,
		room2: room2,
	}
	
	return &door
}