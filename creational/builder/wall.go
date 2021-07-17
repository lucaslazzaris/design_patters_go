package main

type WallInterface interface {
	enter()
}

type Wall struct {
}

func (wall *Wall) enter() {
}