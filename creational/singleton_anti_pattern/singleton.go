package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Singleton struct {
}

var singleInstance *Singleton

func getInstance() *Singleton {
    if singleInstance == nil {
        lock.Lock()
        defer lock.Unlock()
        if singleInstance == nil {
            fmt.Println("Creating single instance now.")
            singleInstance = &Singleton{}
        } else {
            fmt.Println("Single instance already created.")
        }
    } else {
        fmt.Println("Single instance already created.")
    }

    return singleInstance
}

func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}
}