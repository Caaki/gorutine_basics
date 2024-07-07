package main

import (
	"fmt"
	"math/rand"
)

// Za ovaj primer je bitno primetiti i zapamtiti da se nit zaustavlja posto ceka podatke od kanala c
// Zato nije potrebno implementirati wait grupe
func getRandomNumber(c chan int32) {
	theRandomNumber := rand.Int31n(1000)
	fmt.Printf("The random number is: %d\n", theRandomNumber)

	c <- theRandomNumber

}

func main() {
	c := make(chan int32)
	go getRandomNumber(c)

	//output := <-c
	fmt.Printf("%d", <-c)

}
