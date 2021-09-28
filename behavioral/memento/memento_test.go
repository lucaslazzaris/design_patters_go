package main

import "testing"

func TestMemento(t *testing.T){
	state := new(ConstraintSolverMemento)
	command := MoveCommand{state: state, delta: Point{}, target: new(Graphic)}

	executeActual := command.execute()
	unexecuteActual := command.unexecute()

	executeExpected := "Moving direction 1 Solving"
	unexecuteExpected := "Moving direction -1 Solving"

	if executeActual != executeExpected {
		t.Errorf("Actual: %s Expected: %s",  executeActual, executeExpected)
	}

	if unexecuteActual != unexecuteExpected {
		t.Errorf("Actual: %s Expected: %s",  unexecuteActual, unexecuteExpected)
	}
}