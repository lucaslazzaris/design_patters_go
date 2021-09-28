package main

type Component struct {
}

type Compositor interface {
	compose() string
}

type SimpleCompositor struct {
}

func (compositor *SimpleCompositor) compose() string {
	return "Simple Composing"
}

type LatexCompositor struct {
}

func (compositor *LatexCompositor) compose() string {
	return "TeX Composing"
}

type ArrayCompositor struct {
}

func (compositor *ArrayCompositor) compose() string {
	return "Array Composing"
}

type Composition struct {
	compositor Compositor
}

func (composition *Composition) repair() string {
	return composition.compositor.compose()
}
