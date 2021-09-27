package main

import "testing"

func TestFlyweight(t *testing.T){
	glyphFactory := new(GlyphFactory)
	glyphFactory.characters = make(map[string]Character)

	glyphFactory.createCharacter("a")
	glyphFactory.createCharacter("b")

	if glyphFactory.characters["a"].charcode != "a" {
		t.Errorf("Actual: %s Expected: %s", glyphFactory.characters["a"].charcode, "a")
	}
	if glyphFactory.characters["b"].charcode != "b" {
		t.Errorf("Actual: %s Expected: %s", glyphFactory.characters["b"].charcode, "b")
	}
}