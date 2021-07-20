package main

func main() {
	mazeFactory := new(EnchantedMazeFactory)
	mazeGame := new(MazeGame)

	mazeGame.createMaze(mazeFactory)

	standardMazeFactory := new(MazeFactory)
	standardMazeGame := new(MazeGame)

	standardMazeGame.createMaze(standardMazeFactory)
}