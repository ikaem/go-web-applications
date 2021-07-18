package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstname() string {
	return m.FirstName

}

func main() {

	var myVar myStruct
	myVar.FirstName = "Karlo"

	myVar2 := myStruct{
		FirstName: "Mike",
	}

	log.Println("My var is set to", myVar)   // My var is set to {Karlo}
	log.Println("My var2 is set to", myVar2) // My var2 is set to {Mike}

	log.Println(myVar.printFirstname()) // karlo

}
