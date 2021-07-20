package main

type EnchantedMazeFactory struct {
}

func (mazeFactory *EnchantedMazeFactory) makeMaze() (*Maze) {
	var maze Maze
	maze.rooms = make(map[int]Spacer)
	return &maze
}

func (mazeFactory *EnchantedMazeFactory) makeWall() (*Wall) {
	return new(Wall)
}

func (mazeFactory *EnchantedMazeFactory) makeRoom(n int) (Spacer){
	room := EnchantedRoom {
		roomNo: n,
		spell: "Alohomora",
	}
	return &room
}

func (mazeFactory *EnchantedMazeFactory) makeDoor(r1 Spacer, r2 Spacer) (Blocker) {
	room1 := r1.(*EnchantedRoom);
	room2 := r2.(*EnchantedRoom);
	
	door := DoorNeedingSpell {
		room1: room1,
		room2: room2,
	}
	return &door
}