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
	wg := sync.WaitGroup{}

	wg.Add(1)
	go processMessages(&wg)

	wg.Add(1)
	go addMessageForRecipiants("Ovo je proba", []string{"Pera", "Mika"}, beacon.signal, 1, &wg)
	wg.Add(1)
	go addMessageForRecipiants("Ovo je proba2", []string{"Pera", "Mika"}, beacon.signal, 7, &wg)

	wg.Wait()

}

func processMessages(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		beacon.signal.L.Lock()
		for len(notReadMessages) == 0 {
			beacon.signal.Wait()
		}

		sendMessageToRecipiants(beacon.signal)
		beacon.signal.L.Unlock()
	}
}

func addMessageForRecipiants(message string, recipients []string, cond *sync.Cond, sec int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(sec) * time.Second)
	cond.L.Lock()
	for _, value := range recipients {
		if len(notReadMessages[value]) == 0 {
			notReadMessages[value] = make([]string, 0)
		}
		notReadMessages[value] = append(notReadMessages[value], message)
	}
	cond.Signal()
	cond.L.Unlock()
}

func sendMessageToRecipiants(cond *sync.Cond) {
	fmt.Println(len(notReadMessages))
	for recipient, messages := range notReadMessages {
		for _, message := range messages {
			fmt.Println("Message for: %s [%s]", recipient, message)
		}
	}

	notReadMessages = make(map[string][]string)
}
