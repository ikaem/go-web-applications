// main_test.go

package main

import (
	"log"
	"testing"
)

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.00, 25.00, 4.00, false},
	{name: "invalid-data", dividend: 100.00, divisor: 0.00, expected: 0.00, isErr: true},
	{name: "expect-5", dividend: 100.00, divisor: 25.00, expected: 5.00, isErr: false},
}

func TestDivide(t *testing.T) {

	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)

		log.Printf("this is got: %v, and this is err: %v", got, err)

		// first one we expect an error, via the property on the struct
		if tt.isErr {
			// but here, there is no error
			if err == nil {
				t.Error("expected an error but did not get one", err)
			}
			// here we do not expect an error
		} else {
			// but we did get an error
			if err != nil {
				t.Error("expected not an error, but did get one", err)
			}
		}

		if got != tt.expected {
			t.Errorf("extpected %f but got %f", tt.expected, got)
		}

	}

}

// func TestDivide(t *testing.T) {
// 	_, err := divide(10.0, 1.0)

// 	if err != nil {
// 		t.Error(("got an error when we should not have"))

// 	}

// 	_, err2 := divide(10.0, 0)

// 	if err2 == nil {
// 		t.Error(("did not get an error when we should have"))

// 	}
// }
