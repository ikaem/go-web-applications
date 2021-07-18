package main

import "fmt"

func main() {

	var myString string
	myString = "green"

	fmt.Println("My string is", myString)

	changeUsingPointer(&myString)
	fmt.Println("My string is", myString)

}

func changeUsingPointer(s *string) {

	fmt.Println("this is the pointer address", s) // this is the pointer address 0xc000088230

	newValue := "Red"

	// we set value at this locaion to Red
	*s = newValue
}
