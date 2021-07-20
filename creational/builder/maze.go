package main

type Maze struct {
	rooms map[int]Spacer
}

func (maze *Maze) addRoom(room Spacer) {
	maze.rooms[room.getRoomNo()] = room
}

func (maze *Maze) roomNo(roomNo int) (Spacer, bool) {
	room, ok := maze.rooms[roomNo]
	return room, ok
}