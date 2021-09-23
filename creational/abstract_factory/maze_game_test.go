package main

import "testing"

func TestMazeGame(t *testing.T) {
	t.Run("StandardMazeFactory", func(t *testing.T) {
		standardMazeFactory := new(MazeFactory)
		mazeGame := new(MazeGame)
		maze :=	mazeGame.createMaze(standardMazeFactory)

		if len(maze.rooms) != 2 {
			t.Errorf("actual %d, expected %d", len(maze.rooms), 2)
		}
		
		_, isDoor := maze.rooms[1].getSide(east).(*Door)
		_, isWall := maze.rooms[2].getSide(east).(*Wall)
		if !isDoor {
			t.Errorf("actual %t, expected %t", isDoor, true)
		}

		if !isWall {
			t.Errorf("actual %t, expected %t", isWall, true)
		}
	})

	t.Run("EnchantedMazeFactory", func(t *testing.T) {
		enchantedMazeFactory := new(EnchantedMazeFactory)
		mazeGame := new(MazeGame)
		maze :=	mazeGame.createMaze(enchantedMazeFactory)

		if len(maze.rooms) != 2 {
			t.Errorf("actual %d, expected %d", len(maze.rooms), 2)
		}
		
		_, isDoor := maze.rooms[1].getSide(east).(*DoorNeedingSpell)
		_, isWall := maze.rooms[2].getSide(east).(*Wall)
		if !isDoor {
			t.Errorf("actual %t, expected %t", isDoor, true)
		}

		if !isWall {
			t.Errorf("actual %t, expected %t", isWall, true)
		}

		enchantedRoom, _ := maze.rooms[1].(*EnchantedRoom)
		if enchantedRoom.spell != "Alohomora" {
			t.Errorf("actual %s, expected %s", enchantedRoom.spell, "Alohomora")
		}
	})	
}