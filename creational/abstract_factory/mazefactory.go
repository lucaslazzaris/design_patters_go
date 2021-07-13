package main

type MazeFactoryInterface interface {
	MakeMaze() *Maze
	MakeWall() *Wall
	MakeRoom(int) Spacer 
	MakeDoor(Spacer, Spacer) Blocker
}

type MazeFactory struct {
}

func (mazeFactory *MazeFactory) MakeMaze() (*Maze) {
	var maze Maze
	maze.rooms = make(map[int]Spacer)
	return &maze
}

func (mazeFactory *MazeFactory) MakeWall() (*Wall) {
	return new(Wall)
}

func (mazeFactory *MazeFactory) MakeRoom(n int) (Spacer){
	room := Room {
		roomNo: n,
	}
	return &room
}

func (mazeFactory *MazeFactory) MakeDoor(r1 Spacer, r2 Spacer) (Blocker) {
	room1 := r1.(*Room);
	room2 := r2.(*Room);
	
	door := Door {
		room1: room1,
		room2: room2,
	}
	
	return &door
}