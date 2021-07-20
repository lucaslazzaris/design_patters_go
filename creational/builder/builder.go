package main

func main() {
	mazeBuilder := new(CountingMazeBuilder)
	mazeGame := new(MazeGame)

	mazeGame.createMaze(mazeBuilder)

	standardMazeBuilder := new(StandardMazeBuilder)
	standardMazeGame := new(MazeGame)

	standardMazeGame.createMaze(standardMazeBuilder)
}