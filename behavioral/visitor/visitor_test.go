package main

import "testing"

func TestVisitor(t *testing.T){
	disk := new(FloppyDisk)
	chassis := new(Chassis)
	card := new(Card)
	bus := new(Bus)

	visitor := PricingVisitor{total: 0}
	disk.accept(&visitor)
	chassis.accept(&visitor)
	card.accept(&visitor)
	bus.accept(&visitor)

	actualTotal := visitor.total
	expectTotal := float32(19.0)

	if actualTotal != expectTotal {
		t.Errorf("Actual: %f Expected: %f",  actualTotal, expectTotal)
	}
}