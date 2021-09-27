package main

import "testing"

func TestFacade(t *testing.T){
	compiler := Compiler{}

	actualResponse := compiler.compile("input", "output")
	expectedResponse := "Testing this component requires a lot of not implemented integrations"

	if actualResponse != expectedResponse {
		t.Errorf("Actual: %s Expected: %s", actualResponse, expectedResponse)

	}

}