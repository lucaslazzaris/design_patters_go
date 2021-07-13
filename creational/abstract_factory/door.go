package main

import (
	"errors"
)

type Blocker interface {
	OtherSideRoom(Spacer) (Spacer, error)
	enter()
}

type Door struct {
	room1 *Room
	room2 *Room
	isOpen bool
}

func (door *Door) OtherSideRoom(room Spacer) (Spacer, error) {
	if(room == door.room1){
		return door.room2, nil
	} else if (room == door.room2) {
		return room, nil
	}

	return nil, errors.New("Door is not cointained in this room")
}

func (door *Door) enter() {
}

type DoorNeedingSpell struct {
	room1 *EnchantedRoom
	room2 *EnchantedRoom
	isOpen bool
}

func (door *DoorNeedingSpell) OtherSideRoom(room Spacer) (Spacer, error) {
	if(room == door.room1){
		return door.room2, nil
	} else if (room == door.room2) {
		return room, nil
	}

	return nil, errors.New("Door is not cointained in this room")
}

func (door *DoorNeedingSpell) enter() {
}