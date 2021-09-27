package main

import "testing"

func TestComposite(t *testing.T){
		cabinet := new(Cabinet)
		chassis := new(Chassis)
		disk := new(FloppyDisk)

		cabinet.add(chassis)
		chassis.add(disk)

		actualDiscontPrice := cabinet.discountPrice().value
		actualNetPrice := cabinet.netPrice().value
		actualPower := cabinet.power().value


		expectedDiscontPrice := 127
		expectedNetPrice := 130
		expectedPower := 93

		if actualDiscontPrice != expectedDiscontPrice {
			t.Errorf("Actual: %d Expected: %d", actualDiscontPrice, expectedDiscontPrice)
		}

		if actualNetPrice != expectedNetPrice {
			t.Errorf("Actual: %d Expected: %d", actualNetPrice, expectedNetPrice)
		}

		if actualPower != expectedPower {
			t.Errorf("Actual: %d Expected: %d", actualPower, expectedPower)
		}
}