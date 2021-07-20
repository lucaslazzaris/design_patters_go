package main

type Blocker interface {
	enter()
}

type Door struct {
	room1 *Room
	room2 *Room
	isOpen bool
}

func (door *Door) enter() {
}

type DoorNeedingSpell struct {
	room1 *EnchantedRoom
	room2 *EnchantedRoom
	isOpen bool
}

func (door *DoorNeedingSpell) enter() {
}