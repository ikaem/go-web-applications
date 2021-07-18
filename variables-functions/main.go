package main

import "fmt"

// var notNeededToBeUsed = 6

func main() {

	fmt.Println("Hello world")

	var whatToSay string
	var i int

	whatToSay = "Goodbye cruel world"

	fmt.Println(whatToSay)

	i = 62

	fmt.Println("i is set to", i)

	whatWasSaid, anotherSaid := saySomething()

	fmt.Println("The function returned", whatWasSaid, anotherSaid)

}

func saySomething() (string, string) {
	return "Something", "Second somehting"
}
