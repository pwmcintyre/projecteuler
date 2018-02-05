package main

import (
	"reflect"
	"testing"
)

var examples = []struct {
	input  int
	answer int
}{
	{3 + 4 + 5, 3 * 4 * 5},
	{1000, 31875000},
}

func TestRun(t *testing.T) {
	for _, ex := range examples {
		_, result, _ := Run(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
