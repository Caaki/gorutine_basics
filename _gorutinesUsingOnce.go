package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var completed = false

func main() {

	wg := sync.WaitGroup{}
	var once sync.Once
	for i := range 500 {
		wg.Add(1)
		if tryFinishingTask() {
			fmt.Printf("Task was completed by: %d\n", i)
			once.Do(taskCompleted)
		}
		wg.Done()
	}
	wg.Wait()
	fmt.Println(completed)

}
func tryFinishingTask() bool {

	if value := rand.Intn(100); value == 77 {
		return true
	}
	return false
}

func taskCompleted() {
	completed = true
	fmt.Println("TASK COMPLETED FUNCTION WAS CALLED")
}
