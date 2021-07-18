package main

import (
	"fmt"
	"time"
)

// this type is inferred
// var s = "seven"

// var birthDate time.Time

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	birthDate   time.Time
}

func main() {

	user := User{
		FirstName:   "Karlo",
		LastName:    "Marinovic",
		PhoneNumber: "999",
		Age:         35,
		// birthDate: time.Time("2021-09-21"),
	}

	fmt.Println(user)
	/*
		{Karlo Marinovic 999 35 {0 0 <nil>}}
	*/
	fmt.Println(user.Age)       // 35
	fmt.Println(user.birthDate) // 0001-01-01 00:00:00 +0000 UTC

	/* 	s := 3 // this is a new s variable - we are declaring new variable

	   	log.Println(s) // 2021/07/18 14:52:35 3 */

}

/* func saySomething(s string) (string, string) {
	return s, "World"
}
*/

func Whatever() {
	fmt.Println("hello")
}
