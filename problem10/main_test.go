package main

import (
	"reflect"
	"testing"
)

var examples = []struct {
	input  int
	answer int
}{
	{10, 17},
	{100, 1060},
	{2000000, 142913828922},
}

func TestRun(t *testing.T) {
	for _, ex := range examples {
		result := Run(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
