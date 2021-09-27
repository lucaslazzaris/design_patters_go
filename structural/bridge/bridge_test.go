package main

import "testing"

func TestBridge(t *testing.T){
	applicationWindow := ApplicationWindow{
		view: new(View),
		windowImp: new(ApplicationWindowImp),
	}

	point1 := Point{coord: &Coord{x: 1, y: 1}}
	point2 := Point{coord: &Coord{x: 2, y: 2}}

	actual := applicationWindow.drawRect(point1, point2)
	expected := "Rect on imp 1.0 1.0 2.0 2.0"
	if actual != expected {
		t.Errorf("Actual: %s Expected: %s", actual, expected)
	}
}