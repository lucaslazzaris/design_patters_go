package main

import "testing"

func TestStrategy(t *testing.T){
	composition := Composition{compositor: new(LatexCompositor)}
	latexActual := composition.repair()

	composition.compositor = new(SimpleCompositor)
	simpleActual := composition.repair()

	composition.compositor = new(ArrayCompositor)
	arrayActual := composition.repair()

	latexExpected := "TeX Composing"
	simpleExpected := "Simple Composing"
	arrayExpected := "Array Composing"


	if latexActual != latexExpected {
		t.Errorf("Actual: %s Expected: %s",  latexActual, latexExpected)
	}

	if simpleActual != simpleExpected {
		t.Errorf("Actual: %s Expected: %s",  simpleActual, simpleExpected)
	}

	if arrayActual != arrayExpected {
		t.Errorf("Actual: %s Expected: %s",  arrayActual, arrayExpected)
	}
}