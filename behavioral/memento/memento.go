package main

import "fmt"

type Graphic struct {
}

func (graphic *Graphic) move(point Point, direction int) string {
	return fmt.Sprintf("Moving direction %d", direction)
}

type Point struct {
}

type ConstraintSolver struct {
	memento *ConstraintSolverMemento
}

func (constraintSolver *ConstraintSolver) createMemento() *ConstraintSolverMemento {
	return &ConstraintSolverMemento{}
}

func (constraintSolver *ConstraintSolver) setMemento(memento *ConstraintSolverMemento) {
	constraintSolver.memento = memento
}

func (constraintSolver ConstraintSolver) solve() string {
	return " Solving"
}

type ConstraintSolverMemento struct {
}

type MoveCommand struct {
	state *ConstraintSolverMemento
	delta Point
	target *Graphic
}

func (command *MoveCommand) execute() string {
	solver := new(ConstraintSolver)
	command.state = solver.createMemento()
	moveText := command.target.move(command.delta, 1)
	return moveText + solver.solve()
}

func (command *MoveCommand) unexecute() string {
	solver := new(ConstraintSolver)
	moveText := command.target.move(command.delta, -1)
	solver.setMemento(command.state)
	return moveText + solver.solve()
}

func main() { 
	state := new(ConstraintSolverMemento)
	command := MoveCommand{state: state, delta: Point{}, target: new(Graphic)}

	command.execute()
	command.unexecute()
}
