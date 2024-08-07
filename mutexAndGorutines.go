package main

import (
	"fmt"
	"sync"
	"time"
)

var container = Container{mux: sync.RWMutex{}, value: 0}

type Container struct {
	mux   sync.RWMutex
	value int
}

func main() {

	for range 1000 {
		go addToMutex()
	}

	time.Sleep(3 * time.Second)

	fmt.Println(container.value)

}

func addToMutex() {
	container.mux.Lock()
	container.value++
	container.mux.Unlock()
}
