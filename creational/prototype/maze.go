package main

type Maze struct {
	rooms map[int]Spacer
}

func (maze *Maze) addRoom(room Spacer) {
	maze.rooms[room.getRoomNo()] = room
}

func (maze *Maze) clone() *Maze {
	newMaze := new(Maze)
	newMaze.rooms = make(map[int]Spacer)
	return newMaze
}