package main

import (
	"fmt"
	"sync"
)

func main() {

	enemies := []string{"John", "Tom", "Stan"}

	waitingGroup := sync.WaitGroup{}
	attack(enemies, &waitingGroup)
	waitingGroup.Wait()

}

func attack(targets []string, waitGroup *sync.WaitGroup) {
	waitGroup.Add(1)
	for enemy := range targets {
		fmt.Println("Target attacked: ", enemy)
	}
	waitGroup.Done()
}
