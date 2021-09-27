package main

import "testing"

func TestChainOfResponsability(t *testing.T){
	t.Run("Component handling help", func(t *testing.T) {
		application := newApplication(application_topic)
		dialog := newDialog(application, print_topic)
		button := newButton(dialog, paper_orientation_topic)
		
		actualButtonHelp := button.handleHelp()
		expectedButtonHelp := "Button help"

		actualDialogHelp := dialog.handleHelp()
		expectedDialogHelp := "Dialog help"

		if actualButtonHelp != expectedButtonHelp {
			t.Errorf("Actual: %s Expected: %s", actualButtonHelp, expectedButtonHelp)
		}

		if actualDialogHelp != expectedDialogHelp {
			t.Errorf("Actual: %s Expected: %s", actualDialogHelp, expectedDialogHelp)
		}
	})

	t.Run("Chain of command handling help", func(t *testing.T) {
		application := newApplication(application_topic)
		dialog := newDialog(application, application_topic)
	
		actualDialogHelp := dialog.handleHelp()
		expectedDialogHelp := "Application help"
	
		if actualDialogHelp != expectedDialogHelp {
			t.Errorf("Actual: %s Expected: %s", actualDialogHelp, expectedDialogHelp)
		}
	})
	

}