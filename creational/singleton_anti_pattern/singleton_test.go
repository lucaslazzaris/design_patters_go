package main

import "testing"

func TestSingleton(t *testing.T) {
	firstInstance := getInstance()

	if firstInstance.NumberOfCreations() != 1 {
		t.Error("expected just one number of creations")
	}

	secondInstance := getInstance()
	if firstInstance != secondInstance {
		t.Error("expected same instance")
	}

	thirdInstance := getInstance()
	if thirdInstance.NumberOfCalls() != 3 {
		t.Errorf("expected three calls not %d", thirdInstance.NumberOfCalls())
	}

	if firstInstance.NumberOfCreations() != 1 {
		t.Error("expected just one number of creations")
	}
}