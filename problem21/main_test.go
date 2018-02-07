package main

import (
	"reflect"
	"testing"
)

var examples = []struct {
	input  int
	answer int
}{
	{200, 0},
	{1000, 504},
	{2000, 2898},
	{10000, 31626},
}

func TestRun(t *testing.T) {
	for _, ex := range examples {
		result := Run(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
