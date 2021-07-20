package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {

	dog := Dog{
		Name:  "Doggy",
		Breed: "Some",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Joe",
		Color:         "black",
		NumberOfTeeth: 11,
	}

	PrintInfo(&gorilla)

}

func PrintInfo(a Animal) {

	fmt.Println("This animal says ", a.Says(), "and has ", a.NumberOfLegs(), " legs")

}

func (d *Gorilla) Says() string {
	return "Something woof"
}

func (d *Gorilla) NumberOfLegs() int {
	return 4
}

func (d *Dog) Says() string {
	return "Something woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}
