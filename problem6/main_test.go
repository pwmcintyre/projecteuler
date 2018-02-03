package problem6

import (
	"reflect"
	"testing"
)

var examples = []struct {
	input        int
	SumOfSquares int
	SquareOfSum  int
	difference   int
}{
	{1, 1, 1, 0},
	{2, 5, 9, 4},
	{10, 385, 3025, 2640},
	{100, 338350, 25502500, 25164150},
}

func TestRun(t *testing.T) {

	for _, ex := range examples {
		result := Run(ex.input)
		if !reflect.DeepEqual(ex.difference, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.SquareOfSum, ex.input)
		}
	}
}

func TestSumOfSquares(t *testing.T) {

	for _, ex := range examples {
		result := SumOfSquares(ex.input)
		if !reflect.DeepEqual(ex.SumOfSquares, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.SumOfSquares, ex.input)
		}
	}
}

func TestSquareOfSum(t *testing.T) {

	for _, ex := range examples {
		result := SquareOfSum(ex.input)
		if !reflect.DeepEqual(ex.SquareOfSum, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.SquareOfSum, ex.input)
		}
	}
}
