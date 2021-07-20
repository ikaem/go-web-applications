package main

import (
	"errors"
	"log"
)

func main() {

	result, err := divide(24.00, 38.2)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("result is", result)

	result2, err := divide(24.00, 0)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("result is", result2)
}

func divide(x, y float32) (float32, error) {

	var result float32
	log.Println("this is y", y)

	if y == 0 {
		return result, errors.New("cannot divide by zero")
	}

	result = x / y
	return result, nil
}
