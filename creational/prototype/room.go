package main

const (
	south = 0
	north = 1
	east = 2
	west = 3
)

type Spacer interface {
	getSide(direction int) MapSite
	setSide(direction int, mapSite MapSite) 
	enter()
	getRoomNo() int
	getSides() [4]MapSite
}

type Room struct {
	roomNo int
	sides [4]MapSite
}

func (room *Room) getSide(direction int) MapSite {
	return room.sides[direction]
}

func (room *Room) setSide(direction int, mapSite MapSite) {
	room.sides[direction] = mapSite
}

func (room *Room) enter() {
}

func (room *Room) getRoomNo() int {
	return room.roomNo
}

func (room *Room) getSides() [4]MapSite {
	return room.sides
}

type EnchantedRoom struct {
	roomNo int
	sides [4]MapSite
	spell string
}

func (room *EnchantedRoom) getSide(direction int) MapSite{
	return room.sides[direction]
}

func (room *EnchantedRoom) setSide(direction int, mapSite MapSite){
	room.sides[direction] = mapSite
}

func (room *EnchantedRoom) enter() {
}

func (room *EnchantedRoom) getRoomNo() int {
	return room.roomNo
}

func (room *EnchantedRoom) getSides() [4]MapSite{
	return room.sides
}