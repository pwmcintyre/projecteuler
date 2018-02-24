package main

import (
	"fmt"
	"testing"
)

type Example struct {
	input  Input
	answer int
}

func TestProblem1(t *testing.T) {

	examples := []Example{
		{Input{10, 3, 5}, 23},
		{Input{1000, 3, 5}, 233168},
	}

	fmt.Print(examples)

	for _, ex := range examples {
		result := Run(ex.input)
		if result != ex.answer {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
