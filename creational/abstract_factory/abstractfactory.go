package main

func main() {
	mazeFactory := new(MazeFactory)
	mazeGame := new(MazeGame)

	mazeGame.CreateMaze(mazeFactory)
}