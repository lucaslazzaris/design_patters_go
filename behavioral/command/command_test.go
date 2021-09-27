package main

import "testing"

func TestCommand(t *testing.T){
	application := new(Application)

	application.documents = make([]*Document, 0)

	openCommand := newOpenCommand(application)
	openCommand.execute()

	expectedFileName := "UserData.pdf"
	actualFileName := application.documents[0].name

	expectedLength := 1
	actualLength := len(application.documents)

	if expectedFileName != actualFileName {
		t.Errorf("Actual: %s Expected: %s", expectedFileName, actualFileName)
	}

	if expectedLength != actualLength {
		t.Errorf("Actual: %d Expected: %d", expectedLength, actualLength)
	}
}