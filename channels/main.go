package main

import (
	"log"

	"github.com/ikaem/channels-learn/helpers"
)

const numPool = 10

func CalculateValue(intChan chan int) {

	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber

}

func main() {

	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)
	randNum := <-intChan

	log.Println(randNum) // 1 -> always 1

}
