package main

func main() {
	standardMazeFactory := newStandardFactory(
		new(Maze),
		new(Wall),
		new(Room),
		new(Door),
	)
	mazeGame := new(MazeGame)

	mazeGame.createMaze(standardMazeFactory)

	EnchantedMazeFactory := newStandardFactory(
		new(Maze),
		new(Wall),
		new(EnchantedRoom),
		new(DoorNeedingSpell),
	)

	standardMazeGame := new(MazeGame)

	standardMazeGame.createMaze(EnchantedMazeFactory)
}