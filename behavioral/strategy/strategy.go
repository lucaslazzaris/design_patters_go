package main

import "fmt"

type Component struct {
}

type Compositor interface {
	compose()
}

type SimpleCompositor struct {
}

func (compositor *SimpleCompositor) compose() {
	fmt.Println("Simple Composing")
}

type LatexCompositor struct {
}

func (compositor *LatexCompositor) compose() {
	fmt.Println("TeX Composing")
}

type ArrayCompositor struct {
}

func (compositor *ArrayCompositor) compose() {
	fmt.Println("Array Composing")
}

type Composition struct {
	compositor Compositor
}

func (composition *Composition) repair(){
	composition.compositor.compose()
}

func main() {
	composition := Composition{compositor: new(LatexCompositor)}
	composition.repair()

	composition.compositor = new(SimpleCompositor)
	composition.repair()

	composition.compositor = new(ArrayCompositor)
	composition.repair()
}