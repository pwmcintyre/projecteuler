package main

import (
	"reflect"
	"testing"
)

var examples = []struct {
	input  int
	answer int
}{
	{1, 1},
	{2, 2},
	{4, 6},
	{10, 27},
	{100, 648},
}

func TestRun(t *testing.T) {
	for _, ex := range examples {
		result := Run(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
