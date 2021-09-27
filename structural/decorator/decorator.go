package main

import "fmt"

type VisualComponent interface {
	draw() string
	resize() string
}

type TextView struct {
}

func (textView *TextView) draw() string {
	return "Drawing"
}

func (textView *TextView) resize() string {
	return "Resizing"
}

type Decorator interface {
	VisualComponent
}

type Decorate struct {
	component VisualComponent
}

func (decorator *Decorate) draw() string {
	return decorator.component.draw()
}

func (decorator *Decorate) resize() string {
	return decorator.component.resize()
}

type BorderDecorator struct {
	component VisualComponent
	width int
}

func (decorator *BorderDecorator) draw() string {
	draw := decorator.component.draw()
	border := decorator.drawBorder(decorator.width)
	return draw + border
}

func (decorator *BorderDecorator) resize() string {
	return decorator.component.resize()
}

func (decorator *BorderDecorator) drawBorder(width int) string { 
	// do something
	return fmt.Sprintf(" Adding border with %d width", width)
}

type ScrollDecorator struct {
	component VisualComponent
	height int
}

func (decorator *ScrollDecorator) draw() string {
	draw := decorator.component.draw()
	scroll := decorator.scroll(decorator.height)
	return draw + scroll
}

func (decorator *ScrollDecorator) resize() string {
	return decorator.component.resize()
}

func (decorator *ScrollDecorator) scroll(height int) string { 
	// do something
	return fmt.Sprintf(" Scrolling %d pixels", height)
}
