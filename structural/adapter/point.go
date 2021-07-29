package main

type Coord struct {
	x float32
	y float32
}

type Point struct {
	coord *Coord
}

func newPoint(x float32, y float32) Point {
	coord := new(Coord)
	coord.x = x
	coord.y = y
	return Point { coord: coord }
}