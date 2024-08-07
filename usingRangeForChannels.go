package main

import (
	"fmt"
	"math/rand"
)

func main() {

	channel := make(chan string)
	channel2 := make(chan string)
	go getNumbers(channel, "1")
	go getNumbers(channel2, "2")

	for {
		select {
		case message, ok := <-channel:
			if ok {
				fmt.Println(message)
			}
		case message, ok := <-channel2:
			if ok {
				fmt.Println(message)
			}
		default:

		}
		//	if message, ok := <-channel; ok {
		//		fmt.Println(message)
		//	}
		//	break
		//}
	}
}

func getNumbers(channel chan string, message string) {

	numberOfRound := rand.Intn(5) + 1
	fmt.Println(numberOfRound)

	for range numberOfRound {
		channel <- fmt.Sprintf(" Message from channel %s [%d]", message, rand.Intn(10))
	}
	close(channel)
}
