package main

import "testing"

func TestMazePrototypeFactory(t *testing.T) {
	t.Run("StandardMazePrototypeFactory", func(t *testing.T) {
		standardMazeFactory := newStandardFactory(
			new(Maze),
			new(Wall),
			new(Room),
			new(Door),
		)

		newMaze := standardMazeFactory.makeMaze()
		newRoom1 := standardMazeFactory.makeRoom(1)
		newRoom2 := standardMazeFactory.makeRoom(2)
		newDoor := standardMazeFactory.makeDoor(newRoom1, newRoom2)
		newWall := standardMazeFactory.makeWall()

		_, isRoom1 := newRoom1.(*Room)
		_, isRoom2 := newRoom2.(*Room)
		_, isDoor := newDoor.(*Door)

		if newMaze == nil {
			t.Error("maze is nil")
		}	
		if newWall == nil {
			t.Error("wall is nil")
		}	
	
		if !isRoom1 || !isRoom2 {
			t.Error("isRoom is not of type Room")
		}
		
		if !isDoor {
			t.Error("isDoor is not of type Door")
		}
	})

	t.Run("EnchantedMazePrototypeFactory", func(t *testing.T) {
		standardMazeFactory := newStandardFactory(
			new(Maze),
			new(Wall),
			new(EnchantedRoom),
			new(DoorNeedingSpell),
		)

		newMaze := standardMazeFactory.makeMaze()
		newRoom1 := standardMazeFactory.makeRoom(1)
		newRoom2 := standardMazeFactory.makeRoom(2)
		newDoor := standardMazeFactory.makeDoor(newRoom1, newRoom2)
		newWall := standardMazeFactory.makeWall()

		enchantedRoom, isRoom1 := newRoom1.(*EnchantedRoom)
		_, isRoom2 := newRoom2.(*EnchantedRoom)
		_, isDoor := newDoor.(*DoorNeedingSpell)

		if newMaze == nil {
			t.Error("maze is nil")
		}	
		if newWall == nil {
			t.Error("wall is nil")
		}	
	
		if !isRoom1 || !isRoom2 {
			t.Error("isRoom is not of type Room")
		}
		
		if !isDoor {
			t.Error("isDoor is not of type Door")
		}

		if enchantedRoom.spell != "Alohomora" {
			t.Errorf("actual %s, expected %s", enchantedRoom.spell, "Alohomora")
		}
	})
}