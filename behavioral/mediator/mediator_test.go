package main

import "testing"

func TestMediator(t *testing.T){
	director := FontDialogDirector{}
	director.createWidgets()

	cancelButtonChange := director.cancel.changed()
	okButtonChange := director.ok.changed()
	fontListChanged := director.fontList.changed()
	fontNameChanged := director.fontName.changed()

	if cancelButtonChange != "Cancel" {
		t.Errorf("Actual: %s Expected: %s", cancelButtonChange, "Cancel")
	}

	if okButtonChange != "Ok" {
		t.Errorf("Actual: %s Expected: %s", cancelButtonChange, "Ok")
	}

	if fontListChanged != "Font List" {
		t.Errorf("Actual: %s Expected: %s", fontListChanged, "Font List")
	}

	if fontNameChanged != "Font Name" {
		t.Errorf("Actual: %s Expected: %s", fontNameChanged, "Font Name")
	}
}