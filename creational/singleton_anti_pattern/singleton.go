package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Singleton struct {
    calls     int
	creations int
}
func (s *Singleton) NumberOfCalls() int {
	return s.calls
}

func (s *Singleton) NumberOfCreations() int {
	return s.creations
}

var singleInstance *Singleton

func getInstance() *Singleton {
    if singleInstance == nil {
        lock.Lock()
        defer lock.Unlock()
        fmt.Println("Creating single instance now.")
        singleInstance = &Singleton{calls: 0, creations: 0}
        singleInstance.creations++
        singleInstance.calls++
    } else {
        singleInstance.calls++
        fmt.Println("Single instance already created.")
    }

    return singleInstance
}

func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}
}