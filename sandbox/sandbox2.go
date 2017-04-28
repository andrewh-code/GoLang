package main

import (
	"fmt"
	"math/rand"
)

func returnNumber() (randomNumber int) {
	//return a random number
	randomNumber = rand.Intn(99)
	fmt.Println("returning a random number from sandbox2.go, function returnNumber()")

	return
}
