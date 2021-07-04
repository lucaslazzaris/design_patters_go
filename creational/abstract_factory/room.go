package main

const (
	south = 0
	north = 1
	east = 2
	west = 3
)

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