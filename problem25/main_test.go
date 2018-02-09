package main

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {

	var examples = []struct {
		input  int
		answer int
	}{
		{1, 1}, // 1
		{2, 7}, // 1, 1, 2, 3, 5, 8, 13
		{3, 12},
		{1000, 4782},
	}

	for _, ex := range examples {
		result := Run(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
