package main

import "testing"

func TemplateTest(t *testing.T){
	view := View{template: new(ScreenView)}

	actualDisplay := view.display()
	expectedDisplay := "Really Focused Display with Screen View Focus Resetted"


	if actualDisplay != expectedDisplay {
		t.Errorf("Actual: %s Expected: %s",  actualDisplay, expectedDisplay)
	}
}