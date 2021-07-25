package main

type Wall struct {
}

func (wall *Wall) enter() {
}

func (wall *Wall) clone() *Wall{
	return new(Wall)
}