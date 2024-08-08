package main

import (
	"fmt"
	"sync"
	"time"
)

type ConditionWithCond struct {
	cond   bool
	signal *sync.Cond
}

var beacon = ConditionWithCond{cond: false, signal: sync.NewCond(&sync.Mutex{})}
var notReadMessages = make(map[string][]string)
var channel = make(chan bool)

func main() {
	defer close(channel)
	//wg := sync.WaitGroup{}

	go addMessageForRecipiants("Ovo je proba", []string{"Pera", "Mika"}, beacon.signal, 7)
	go addMessageForRecipiants("Ovo je proba2", []string{"Pera", "Mika"}, beacon.signal, 1)

	for {
		select {
		case message, ok := <-channel:
			if ok && message == true {
				sendMessageToRecipiants(beacon.signal)
			}
			if !ok {
				close(channel)
				break
			}
		}
	}

}

func addMessageForRecipiants(message string, recipients []string, cond *sync.Cond, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	cond.L.Lock()
	for _, value := range recipients {
		if len(notReadMessages[value]) == 0 {
			notReadMessages[value] = make([]string, 0)
		}
		notReadMessages[value] = append(notReadMessages[value], message)
	}
	cond.L.Unlock()
	channel <- true
	cond.Signal()
}

func sendMessageToRecipiants(cond *sync.Cond) {

	cond.L.Lock()
	for recipient, messages := range notReadMessages {
		for _, message := range messages {
			fmt.Println("Message for: %s [%s]", recipient, message)
		}
	}

	cond.Wait()
	cond.L.Unlock()
	channel <- false
}
