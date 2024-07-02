package main

import (
	"fmt"
	"time"
)

// Ovaj primer pokazuje kako mozemo da postavimo kanal i proveravamo da li ima poruku i kada se pojavi
// poruka odma je ispisemo
func main() {

	kanal := make(chan string, 2)
	kanal <- "Prva poruka"
	kanal <- "Druga poruka"

	go func() {
		time.Sleep(time.Second * 5)
		kanal <- "Treca poruka"
	}()

	sleepTimeLimit := 6
	currentSleepTime := 0
	idle := false
	for !idle {
		select {
		case x, ok := <-kanal:
			if ok {
				fmt.Println(x)
				currentSleepTime = 0
			} else {
				break
			}
		default:
			time.Sleep(time.Second * 1)
			fmt.Printf("Current sleep time is %d \n", currentSleepTime)
			currentSleepTime++
			if currentSleepTime > sleepTimeLimit {
				idle = true
			}
		}
	}

	fmt.Println("Program ended, server was idle for too long")

}
