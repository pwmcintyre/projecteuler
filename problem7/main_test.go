package main

import (
	"reflect"
	"testing"
)

var examples = []struct {
	input  int
	answer int
}{
	{1, 2},
	{6, 13},
	{1000, 7919},
	{2000, 17389},
	// {10001, 104743}, // too slow
}

func TestRun(t *testing.T) {

	for _, ex := range examples {
		result := Run(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
