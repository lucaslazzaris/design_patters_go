package main

type Maze struct {
	rooms map[int]*Room
}

func (maze *Maze) addRoom(room *Room) {
	maze.rooms[room.roomNo] = room
}

func (maze *Maze) RoomNo(roomNo int) *Room {
	return maze.rooms[roomNo]
}