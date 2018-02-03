package functions

import (
	"reflect"
	"testing"
)

func TestGCD(t *testing.T) {

	// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
	examples := []struct {
		input  []int
		answer int
	}{
		{[]int{48, 18}, 6},
		{[]int{10, 20}, 10},
		{[]int{10, 21}, 1},
		{[]int{123, 321}, 3},
	}

	for _, ex := range examples {
		result := GCD(ex.input[0], ex.input[1])
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
