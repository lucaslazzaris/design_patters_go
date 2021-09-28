package main

import "testing"

func TestInterpreter(t *testing.T) {
	context := Context{cont: make(map[VariableExp]bool)}

	x := VariableExp{name: "X"}
	y := VariableExp{name: "Y"}

	expression := OrExp{
		op1: &AndExp{op1: &y, op2: &x},
		op2: &AndExp{op1: &y, op2: &NotExp{op: &x}},
	}

	context.assign(x, false)
	context.assign(y, true)

	actual := expression.evaluate(context)
	expected := true

	if actual != expected {
		t.Errorf("Actual: %t Expected: %t", actual, expected)
	}
}