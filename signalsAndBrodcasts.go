package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ConditionWithCond struct {
	cond   bool
	signal *sync.Cond
}

var ready bool
var beacon = ConditionWithCond{cond: false, signal: sync.NewCond(&sync.Mutex{})}
var notReadMessages = make(map[string][]string)

func main() {

	go func(con ConditionWithCond) {
		for {
			if ready {

				beacon.signal.Wait()
				sendMessageToRecipiants(beacon.signal)
			}
		}
	}(beacon)

	go addMessageForRecipiants("Ovo je proba", []string{"Pera", "Mika"}, beacon.signal)
	go addMessageForRecipiants("Ovo je proba2", []string{"Pera", "Mika"}, beacon.signal)

	time.Sleep(10 * time.Second)

}

func addMessageForRecipiants(message string, recipients []string, cond *sync.Cond) {
	cond.L.Lock()
	for _, value := range recipients {
		if len(notReadMessages[value]) == 0 {
			notReadMessages[value] = make([]string, 0)
		}
		notReadMessages[value] = append(notReadMessages[value], message)
	}
	cond.Wait()
	cond.L.Unlock()
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	cond.Signal()
	ready = true
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

}
