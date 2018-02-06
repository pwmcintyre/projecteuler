package main

import (
	"math/big"
	"reflect"
	"testing"
)

var examples = []struct {
	input  uint
	answer int64
}{
	{1, 2},
	{2, 6},
	{3, 20},
	{4, 70},
	{10, 184756},
	{20, 137846528820},
}

func TestRun(t *testing.T) {
	for _, ex := range examples {
		result := Run(ex.input)
		answer := big.NewInt(ex.answer)
		if !reflect.DeepEqual(answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, answer, ex.input)
		}
	}
}
