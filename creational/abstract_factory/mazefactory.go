package main

type MazeFactory struct {
}

func (mazeFactory *MazeFactory) MakeMaze() (*Maze) {
	var maze Maze
	maze.rooms = make(map[int]*Room)
	return &maze
}

func (mazeFactory *MazeFactory) MakeWall() (*Wall) {
	return new(Wall)
}

func (mazeFactory *MazeFactory) MakeRoom(n int) (*Room){
	room := Room {
		roomNo: n,
	}
	return &room
}

func (mazeFactory *MazeFactory) MakeDoor(r1 *Room, r2 *Room) (*Door) {
	door := Door {
		room1: r1,
		room2: r2,
	}
	return &door
}