package main

import (
	"fmt"
	"math/rand"
)

func main() {

	channel := make(chan int)

	go getNumbers(channel)

	for {
		message, ok := <-channel
		if !ok {
			break
		}
		fmt.Println(message)
	}

}

func getNumbers(channel chan int) {

	for i := 0; i < 3; i++ {
		channel <- rand.Intn(10)
	}
	close(channel)

}
