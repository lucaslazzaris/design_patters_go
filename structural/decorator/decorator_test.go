package main

import "testing"

func TestDecorator(t *testing.T){
	textView := new(TextView)

	scrollDecorator := ScrollDecorator {
		component: textView,
		height: 10,
	}
	borderDecorator := BorderDecorator {
		component: &scrollDecorator,
		width: 8,
	}

	expect := "Drawing Scrolling 10 pixels Adding border with 8 width"
	actual := borderDecorator.draw()

	if expect != actual {
		t.Errorf("Actual: %s Expected: %s", actual, expect)
	}
}