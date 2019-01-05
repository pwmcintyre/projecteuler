package main

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {

	var examples = []struct {
		input  input
		answer answer
	}{
		{input{2, 5, 2, 5}, 15},
		{input{2, 100, 2, 100}, 9183},
	}

	for _, ex := range examples {
		result := Run(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
