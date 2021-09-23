package main

import "testing"

func TestMazeGame(t *testing.T) {
	t.Run("StandardMazeBuilder", func(t *testing.T) {
		mazeBuilder := new(StandardMazeBuilder)
		mazeGame := new(MazeGame)
		mazeGame.createMaze(mazeBuilder)

		maze := mazeBuilder.currentMaze
		
		_, isDoor := maze.rooms[1].getSide(east).(*Door)
		_, isWall := maze.rooms[2].getSide(north).(*Wall)

		if !isDoor {
			t.Errorf("actual %t, expected %t", isDoor, true)
		}

		if !isWall {
			t.Errorf("actual %t, expected %t", isWall, true)
		}
	})
	t.Run("CountingMazeBuilder", func(t *testing.T) {
		mazeBuilder := new(CountingMazeBuilder)
		mazeGame := new(MazeGame)
		mazeGame.createMaze(mazeBuilder)

		if mazeBuilder.rooms != 5 {
			t.Errorf("actual %d, expected %d", mazeBuilder.rooms, 5)
		}

		if mazeBuilder.doors != 4 {
			t.Errorf("actual %d, expected %d", mazeBuilder.doors, 4)
		}
	})
}