package main

import "fmt"

func main() {
	bottomPoint := newPoint(1, 2) 
	topPoint := newPoint(2, 3)
	textView := TextView{bottomLeft: &bottomPoint, topRight: &topPoint}
	textShape := TextShape{textView: &textView}

	point1, point2 := textShape.boundingBox()
	manipulator := textShape.createManipulator()
	fmt.Println(point1.coord.x, point1.coord.y)
	fmt.Println(point2.coord.x, point2.coord.y)
	fmt.Println(textShape.isEmpty())
	fmt.Println(manipulator)
}