package main

type MazePrototypeFactory interface {
	makeMaze() *Maze
	makeRoom(int) Spacer
	makeWall() *Wall
	makeDoor(Spacer, Spacer) Blocker
}

type StandardMazePrototypeFactory struct {
	prototypeMaze *Maze
	prototypeRoom Spacer
	prototypeWall *Wall
	prototypeDoor Blocker
}

func newStandardFactory(maze *Maze, wall *Wall, room Spacer, door Blocker) MazePrototypeFactory {
	return &StandardMazePrototypeFactory{
		maze,
		room,
		wall,
		door,
	}
}

func (factory *StandardMazePrototypeFactory) makeMaze() *Maze {
	return factory.prototypeMaze.clone()
}

func (factory *StandardMazePrototypeFactory) makeRoom(roomNo int) Spacer {
	return factory.prototypeRoom.clone(roomNo)
}

func (factory *StandardMazePrototypeFactory) makeWall() *Wall {
	return factory.prototypeWall.clone()
}

func (factory *StandardMazePrototypeFactory) makeDoor(room1 Spacer, room2 Spacer) Blocker {
	door := factory.prototypeDoor.clone()
	door.initialize(room1, room2) 
	return door
}