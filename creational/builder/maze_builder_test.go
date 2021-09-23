package main

import "testing"

func TestMazeBuilder(t *testing.T) {
	t.Run("StandardMazeBuilder", func(t *testing.T) {
		mazeBuilder := new(StandardMazeBuilder)
		mazeBuilder.buildMaze()
	
		mazeBuilder.buildRoom(1)
		mazeBuilder.buildRoom(2)
	
		mazeBuilder.buildDoor(1, 2)
	
		maze := mazeBuilder.currentMaze
	
		_, isDoor := maze.rooms[1].getSide(east).(*Door)
		_, isWall := maze.rooms[2].getSide(east).(*Wall)
	
		if !isDoor {
			t.Errorf("actual %t, expected %t", isDoor, true)
		}
	
		if !isWall {
			t.Errorf("actual %t, expected %t", isWall, true)
		}
	})
	t.Run("CountingMazeBuilder", func(t *testing.T) {
		mazeBuilder := new(CountingMazeBuilder)
		mazeBuilder.buildMaze()
	
		mazeBuilder.buildRoom(1)
		mazeBuilder.buildRoom(2)
	
		mazeBuilder.buildDoor(1, 2)
	
		if mazeBuilder.doors != 1 {
			t.Errorf("actual %d, expected %d", mazeBuilder.doors, 1)
		}
	
		if mazeBuilder.rooms != 2 {
			t.Errorf("actual %d, expected %d", mazeBuilder.rooms, 2)
		}
	})
}