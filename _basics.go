package main

import (
	"fmt"
	"sync"
	"time"
)

func printWords(words []string, timeInSec uint64) {
	defer waitGroup.Done()
	for _, word := range words {
		fmt.Println(word)
		time.Sleep(time.Second * time.Duration(timeInSec))
	}
}

var waitGroup = sync.WaitGroup{}

func main() {
	start := time.Now()
	words1 := []string{"one", "two", "three"}
	words2 := []string{"one1", "two1", "three1"}
	waitGroup.Add(1)
	go printWords(words1, 1)
	waitGroup.Add(1)
	go printWords(words2, 2)

	waitGroup.Wait()

	fmt.Println("Time from starting is: " + time.Since(start).String())
}
