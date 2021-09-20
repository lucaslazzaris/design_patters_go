package main

import "fmt"

type ViewTemplate interface {
	doDisplay()
}

type View struct {
	template ViewTemplate
}

func (view *View) display() {
	view.setFocus()
	view.template.doDisplay()
	view.resetFocus()
}

func (view *View) setFocus() {
	fmt.Println("Really Focused")
}

func (view *View) resetFocus() {
	fmt.Println("Focus Resetted")
}

type ScreenView struct {
}

func (view *ScreenView) doDisplay() {
	fmt.Println("Display with Screen View")
}


func main() {
	view := View{template: new(ScreenView)}

	view.display()
}