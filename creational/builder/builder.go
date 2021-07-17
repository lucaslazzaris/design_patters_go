package main

func main() {
	mazeBuilder := new(CountingMazeBuilder)
	mazeGame := new(MazeGame)

	mazeGame.CreateMaze(mazeBuilder)
}