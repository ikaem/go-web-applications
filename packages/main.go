// main.go

package main

import (
	"log"

	"github.com/ikaem/packages-learn/helpers"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType

	myVar.TypeName = "some"
}
