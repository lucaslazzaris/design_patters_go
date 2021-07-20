package main

type Maze struct {
	rooms map[int]Spacer
}

func (maze *Maze) addRoom(room Spacer) {
	maze.rooms[room.getRoomNo()] = room
}
