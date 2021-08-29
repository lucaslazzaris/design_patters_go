package main

import "fmt"

type VisualComponent interface {
	draw()
	resize()
}

type TextView struct {
}

func (textView *TextView) draw() {
	fmt.Println("Drawing")
}

func (textView *TextView) resize() {
	fmt.Println("Resizing")
}

type Decorator interface {
	VisualComponent
}

type Decorate struct {
	component VisualComponent
}

func (decorator *Decorate) draw() {
	decorator.component.draw()
}

func (decorator *Decorate) resize() {
	decorator.component.resize()
}

type BorderDecorator struct {
	component VisualComponent
	width int
}

func (decorator *BorderDecorator) draw() {
	decorator.component.draw()
	decorator.drawBorder(decorator.width)
}

func (decorator *BorderDecorator) resize() {
	decorator.component.resize()
}

func (decorator *BorderDecorator) drawBorder(width int) { 
	// do something
	fmt.Println("Adding border")
}

type ScrollDecorator struct {
	component VisualComponent
	height int
}

func (decorator *ScrollDecorator) draw() {
	decorator.component.draw()
	decorator.scroll(decorator.height)
}

func (decorator *ScrollDecorator) resize() {
	decorator.component.resize()
}

func (decorator *ScrollDecorator) scroll(height int) { 
	// do something
	fmt.Println("Scrolling")
}

func main() {
	// this is a dummy implementation

	textView := new(TextView)
	scrollDecorator := ScrollDecorator {
		component: textView,
	}
	borderDecorator := BorderDecorator {
		component: &scrollDecorator,
	}
}