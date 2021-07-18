package main

import (
	"log"
	"sort"
)

func main() {
	// slices

	/* 	var t []int = []int{1, 2, 3, 4}
	   	var tt = []int{1, 2, 3, 4} */

	var sl []int // this is a slice

	sl = append(sl, 3, 4, 5, 6, 2, 0)

	log.Println(sl) //[3]

	sort.Ints(sl)

	log.Println(sl) // 2021/07/18 16:20:47 [0 2 3 4 5 6]

	numbers := []int{1, 2, 3, 4, 5}

	log.Println(numbers[0:2]) // [1,2]
	log.Println(numbers[:2])  // [1,2]
	log.Println(numbers[:2])  // [1,2]

	// maps below

	// var someMap map[string]string = make(map[string]string)

	// myMap := make(map[string]string)
	// myMap["dog"] = "Dog"

	// myMap["dog"] = "Mike"
	// log.Println(myMap["dog"]) // Mike

	/* 	myMap := make(map[string]int)
	   	myMap["first"] = 1

	   	type User struct {
	   		FirstName string
	   	}
	   	myNewMap := make(map[string]User)

	   	myNewMap["me"] = User{
	   		FirstName: "karlo",
	   	}

	   	log.Println(myNewMap) // 2021/07/18 16:10:27 map[me:{karlo}] */

}
