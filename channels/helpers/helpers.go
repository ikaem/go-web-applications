package helpers

import (
	"math/rand"
	"time"
)

type SomeType struct {
	FirstName string
	LastName  string
}

func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)

	return value
}
