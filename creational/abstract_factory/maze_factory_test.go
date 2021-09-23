package main

import "testing"

func TestMazeFactory(t *testing.T) {
	factory := new(MazeFactory)

	maze := factory.makeMaze() 
	wall := factory.makeWall() 
	room1 := factory.makeRoom(1) 
	room2 := factory.makeRoom(2) 
	door := factory.makeDoor(room1, room2) 

	_, isRoom1 := room1.(*Room)
	_, isRoom2 := room2.(*Room)
	_, isDoor := door.(*Door)

	if maze == nil {
		t.Error("maze is nil")
	}	
	if wall == nil {
		t.Error("wall is nil")
	}	

	if !isRoom1 || !isRoom2 {
		t.Error("isRoom is not of type Room")
	}
	
	if !isDoor {
		t.Error("isDoor is not of type Door")
	}
}