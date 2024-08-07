package main

import (
	"fmt"
	"math/rand"
)

func main() {

	channel := make(chan int)
	go getNumbers(channel)
	for {
		if message, ok := <-channel; ok {
			fmt.Println(message)
		}
		break
	}

}

func getNumbers(channel chan int) {

	numberOfRound := rand.Intn(5) + 1
	fmt.Println(numberOfRound)

	for range numberOfRound {
		channel <- rand.Intn(10)
	}
	close(channel)
}
