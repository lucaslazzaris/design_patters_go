package main

type Shaper interface {
	boundingBox() (Point, Point)
	createManipulator() Manipulator
}

type TextShape struct {
	textView *TextView
}

func (textShape TextShape) boundingBox() (point1 Point, point2 Point) {
	bottom, left := textShape.textView.getOrigin()
	width, height := textShape.textView.getExtent()

	bottomLeft := newPoint(bottom, left);
	topRight := newPoint(bottom + height, left + width)
	return bottomLeft, topRight
}

func (textShape TextShape) getOrigin() (float32, float32) {
	return textShape.textView.getOrigin()
}

func (textShape TextShape) isEmpty() bool {
	return textShape.textView.isEmpty()
}

func (textShape TextShape) createManipulator() Manipulator {
	return TextManipulator{textShape: &textShape}
}