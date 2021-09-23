package main

import "testing"

func TestTextView(t *testing.T) {
	bottomPoint := newPoint(1.0, 2.0) 
	topPoint := newPoint(2.0, 4.0)
	textView := TextView{bottomLeft: &bottomPoint, topRight: &topPoint}

	origin_x, origin_y := textView.getOrigin()
	width, height := textView.getExtent()
	isEmpty := textView.isEmpty()


	if origin_x != 1.0 {
		t.Errorf("Actual: %f Expected: %f", origin_x, 1.0)
	}	
	if origin_y != 2.0 {
		t.Errorf("Actual: %f Expected: %f", origin_y, 2.0)
	}	

	if width != 1.0 {
		t.Errorf("Actual: %f Expected: %f", width, 1.0)
	}	

	if height != 2.0 {
		t.Errorf("Actual: %f Expected: %f", height, 2.0)
	}	

	if isEmpty {
		t.Error("Should not be empty")
	}	
}