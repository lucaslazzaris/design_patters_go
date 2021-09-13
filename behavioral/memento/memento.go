package main

import "fmt"

type Graphic struct {
}

func (graphic *Graphic) move(point Point, direction int) {
	fmt.Println("Moving", point, direction)
}

type Point struct {
}

type ConstraintSolver struct {
	memento *ConstraintSolverMemento
}

func (constraintSolver ConstraintSolver) createMemento() *ConstraintSolverMemento {
	return &ConstraintSolverMemento{}
}

func (constraintSolver ConstraintSolver) setMemento(memento *ConstraintSolverMemento) {
	constraintSolver.memento = memento
}

func (constraintSolver ConstraintSolver) solve() {
	fmt.Println("Solving")
}

type ConstraintSolverMemento struct {
}

type MoveCommand struct {
	state *ConstraintSolverMemento
	delta Point
	target *Graphic
}

func (command MoveCommand) execute() {
	solver := new(ConstraintSolver)
	command.state = solver.createMemento()
	command.target.move(command.delta, 1)
	solver.solve()
}

func (command MoveCommand) unexecute() {
	solver := new(ConstraintSolver)
	command.target.move(command.delta, -1)
	solver.setMemento(command.state)
	solver.solve()
}

func main() {
	fmt.Println("Hello world")
}