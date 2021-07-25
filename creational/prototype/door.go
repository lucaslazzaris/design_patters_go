package main

import (
	"errors"
)

type Blocker interface {
	otherSideRoom(Spacer) (Spacer, error)
	initialize(Spacer, Spacer) 
	clone() Blocker
	enter() 
}

type Door struct {
	room1 *Room
	room2 *Room
	isOpen bool
}

func (door *Door) otherSideRoom(room Spacer) (Spacer, error) {
	if(room == door.room1){
		return door.room2, nil
	} else if (room == door.room2) {
		return room, nil
	}

	return nil, errors.New("Door is not cointained in this room")
}

func (door *Door) enter() {
}

func (door *Door) clone() Blocker {
	return new(Door)
}

func (door *Door) initialize(room1 Spacer, room2 Spacer) {
	door.room1, _ = room1.(*Room)
	door.room2, _ = room2.(*Room)
}

type DoorNeedingSpell struct {
	room1 *EnchantedRoom
	room2 *EnchantedRoom
	isOpen bool
}

func (door *DoorNeedingSpell) otherSideRoom(room Spacer) (Spacer, error) {
	if(room == door.room1){
		return door.room2, nil
	} else if (room == door.room2) {
		return room, nil
	}

	return nil, errors.New("Door is not cointained in this room")
}

func (door *DoorNeedingSpell) enter() {
}

func (door *DoorNeedingSpell) clone() Blocker {
	return new(DoorNeedingSpell)
}

func (door *DoorNeedingSpell) initialize(room1 Spacer, room2 Spacer) {
	door.room1, _ = room1.(*EnchantedRoom)
	door.room2, _ = room2.(*EnchantedRoom)
}