package main

import "log"

func main() {

	myVar := "cat"

	switch myVar {
	case "cat", "pool", "ball":
		{
			log.Println("cat")
		}
	case "dog":
		{
			log.Println("dog")
		}
	default:
		{
			log.Println("none of it")
		}
	}

	// var isTrue bool
	// isTrue = true
	// if isTrue {
	// 	log.Println("is true is", isTrue)
	// } else {
	// 	log.Println("is true is", isTrue)
	// }

}
