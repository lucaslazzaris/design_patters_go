package main

import "testing"

func TestState(t *testing.T){
	connection := newTCPConnection()

	actualFirstState := connection.state.getStateName()
	
	connection.changeState(new(TCPEstablished))

	actualSecondState := connection.state.getStateName()

	expectedFirstState := "TCPClosed"
	expectedSecondState := "TCPEstablished"

	
	if actualFirstState != expectedFirstState {
		t.Errorf("Actual: %s Expected: %s",  actualFirstState, expectedFirstState)
	}


	if actualSecondState != expectedSecondState {
		t.Errorf("Actual: %s Expected: %s",  actualSecondState, expectedSecondState)
	}
}