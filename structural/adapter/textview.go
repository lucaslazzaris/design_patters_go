package main

type TextViewer interface {
	getOrigin() (float32, float32)
	getExtent() (float32, float32)
	isEmpty() bool
}

type TextView struct {
	bottomLeft *Point
	topRight *Point
}

func (textView *TextView) getOrigin() (float32, float32){
	return textView.bottomLeft.coord.x, textView.bottomLeft.coord.y
}

func (textView *TextView) getExtent() (float32, float32){
	width := textView.topRight.coord.x - textView.bottomLeft.coord.x
	height := textView.topRight.coord.y - textView.bottomLeft.coord.y
	return width, height 
}

func (textView *TextView) isEmpty() bool {
	return textView.bottomLeft.coord == nil || textView.topRight.coord == nil
}
