package main

import (
	"fmt"
	"sync"
	"time"
)

var ninjas = []string{"ninja1", "ninja2", "ninja3"}
var group = sync.WaitGroup{}

// Tokom rada ovog primera primetio sam da ako se stavim group add unutar gorutine, main ce se odmah izvristi bez cekanja.
// Moguce je je pored jednostavnog dodavanja da defer zatvara resurse mozemo i pisati kompleksne funkcije ako je to potrebno
func main() {
	start := time.Now()

	defer func() {
		fmt.Println(time.Since(start))
	}()

	for _, ninja := range ninjas {
		group.Add(1)
		go func(ninja string) {
			fmt.Println("Ninja \"" + ninja + "\" has been attacked")
			time.Sleep(time.Second * 1)
			group.Done()
		}(ninja)
	}

	group.Wait()

}
