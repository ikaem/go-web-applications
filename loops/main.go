package main

import "log"

func main() {

	type User struct {
		FirstName string
		LastName  string
		Age       int
	}

	me := User{
		FirstName: "Mike",
		LastName:  "Dawson",
		Age:       72,
	}

	var users []User = []User{me, {FirstName: "May", LastName: "Dawson", Age: 62}}

	for i, user := range users {
		log.Println(i, user)

		/*
				$ go run main.go
		2021/07/19 06:52:56 0 {Mike Dawson 72}
		2021/07/19 06:52:56 1 {May Dawson 62}
		*/
	}

	// var firstLine = "meouw"

	// for i, char := range firstLine {
	// 	log.Println(i, char)

	// 	/*
	// 			$ go run main.go
	// 	2021/07/19 06:46:06 0 109
	// 	2021/07/19 06:46:06 1 101
	// 	2021/07/19 06:46:06 2 111
	// 	2021/07/19 06:46:06 3 117
	// 	2021/07/19 06:46:06 4 119
	// 	*/
	// }

	// animals := []string{"dog", "fish", "cow", "cat"}

	// first one is the current interation
	// second one is the element
	// for i, animals := range animals {
	// 	log.Println(i, animals)

	/*
		$ go run main.go
		2021/07/19 06:39:04 0 dog
		2021/07/19 06:39:04 1 fish
		2021/07/19 06:39:04 2 cow
		2021/07/19 06:39:04 3 cat
	*/
	// }

	// for _, animals := range animals {
	// 	log.Println(animals)
	// }

	// for i := 0; i <= 10; i++ {
	// 	log.Println(i)
	// }

	// animals := make(map[string]string)

	// animals["dog"] = "Mike"
	// animals["cat"] = "Lowljoy"
	// animals["cow"] = "Dan"

	// for k, value := range animals {

	// 	log.Println(k, value)

	// 	/*
	// 			$ go run main.go
	// 	2021/07/19 06:43:51 dog Mike
	// 	2021/07/19 06:43:51 cat Lowljoy
	// 	2021/07/19 06:43:51 cow Dan
	// 	*/
	// }

}
