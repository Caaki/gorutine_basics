package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Ovaj primer pokazuje kako mozemo da postavimo kanal i proveravamo da li ima poruku i kada se pojavi
// poruka odma je ispisemo
func main() {
	//channel := make(chan int)
	//go metodaKojaKoristiRange(channel)
	//

	go metodaSaCasom()
	time.Sleep(time.Second * 20)

}

func randomNum(channel chan int, brojPoruka int) {

	for i := 0; i < brojPoruka; i++ {
		channel <- rand.Intn(5)
	}

	close(channel)
}

func metodaKojaKoristiRange(channel chan int) {

	go randomNum(channel, 5)
	//for {
	//	number, ok := <-channel
	//	if !ok {
	//		break
	//	}
	//	fmt.Println(number)
	//}

	for _ = range channel {
		fmt.Println(<-channel)
	}

}

func metodaSaCasom() {

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
