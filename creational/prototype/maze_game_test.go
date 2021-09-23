package main

import "testing"

func TestMazeGame(t *testing.T){
	t.Run("StandardMazePrototypeFactory", func(t *testing.T) {
		standardMazeFactory := newStandardFactory(
			new(Maze),
			new(Wall),
			new(Room),
			new(Door),
		)
		mazeGame := new(MazeGame)
	
		maze := mazeGame.createMaze(standardMazeFactory)

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

	t.Run("EnchantedMazePrototypeFactory", func(t *testing.T) {
		EnchantedMazeFactory := newStandardFactory(
			new(Maze),
			new(Wall),
			new(EnchantedRoom),
			new(DoorNeedingSpell),
		)
	
		standardMazeGame := new(MazeGame)
	
		maze := standardMazeGame.createMaze(EnchantedMazeFactory)

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