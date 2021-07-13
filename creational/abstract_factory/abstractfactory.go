package main

func main() {
	mazeFactory := new(EnchantedMazeFactory)
	mazeGame := new(MazeGame)

	mazeGame.CreateMaze(mazeFactory)
}