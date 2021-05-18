package main

import (
	"errors"
	"fmt"
)

const (
	south = 0
	north = 1
	east = 2
	west = 3
)

type MapSite interface {
	enter()
}

type Room struct {
	roomNo int
	sides [4]MapSite
}

func (room *Room) getSide(direction int) MapSite{
	return room.sides[direction]
}

func (room *Room) setSide(direction int, mapSite MapSite){
	room.sides[direction] = mapSite
}

func (room *Room) enter() {
}

type Wall struct {
}

func (wall *Wall) enter() {
}

type Door struct {
	room1 *Room
	room2 *Room
	isOpen bool
}

func (door *Door) OtherSideRoom(room *Room) (*Room, error) {
	if(room == door.room1){
		return door.room2, nil
	} else if (room == door.room2) {
		return room, nil
	}

	return nil, errors.New("Door is not cointained in this room")
}

func (door *Door) enter() {
}

type Maze struct {
	rooms map[int]*Room
}

func (maze *Maze) addRoom(room *Room) {
	maze.rooms[room.roomNo] = room
}

func (maze *Maze) RoomNo(roomNo int) *Room {
	return maze.rooms[roomNo]
}

func main() {
	maze := new(Maze)
	maze.rooms = make(map[int]*Room)
	r1 := Room{ 
		roomNo: 0,
	}
	r2 := Room{ 
		roomNo: 1,
	}
	door := Door{
		room1: &r1,
		room2: &r2,
	}

	maze.addRoom(&r1)
	maze.addRoom(&r2)

	r1.setSide(north, new(Wall))
	r1.setSide(east, &door)
	r1.setSide(south, new(Wall))
	r1.setSide(west, new(Wall))

	r2.setSide(north, new(Wall))
	r2.setSide(east, new(Wall))
	r2.setSide(south, new(Wall))
	r2.setSide(west, &door)

	fmt.Println(r1.sides)
	fmt.Println(r2.sides)
}