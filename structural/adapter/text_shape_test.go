package main

import "testing"

func TestTextShape(t *testing.T) {
	bottomPoint := newPoint(1.0, 2.0) 
	topPoint := newPoint(2.0, 3.0)
	textView := TextView{bottomLeft: &bottomPoint, topRight: &topPoint}
	textShape := TextShape{textView: &textView}

	point1, point2 := textShape.boundingBox()

	origin_x, origin_y := textShape.getOrigin()
	isEmpty := textShape.isEmpty()


	if origin_x != 1.0 {
		t.Errorf("Actual: %f Expected: %f", origin_x, 1.0)
	}	
	if origin_y != 2.0 {
		t.Errorf("Actual: %f Expected: %f", origin_y, 2.0)
	}	

	if point1.coord.x != 1.0 {
		t.Errorf("Actual: %f Expected: %f", point1.coord.x, 1.0)
	}	

	if point1.coord.y != 2.0 {
		t.Errorf("Actual: %f Expected: %f", point1.coord.y, 2.0)
	}	

	if point2.coord.x != 2.0 {
		t.Errorf("Actual: %f Expected: %f", point2.coord.x, 2.0)
	}	

	if point2.coord.y != 3.0 {
		t.Errorf("Actual: %f Expected: %f", point2.coord.y, 3.0)
	}	

	if isEmpty {
		t.Error("Should not be empty")
	}	
}