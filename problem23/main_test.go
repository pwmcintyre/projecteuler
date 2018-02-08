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
		{1, 0},
		{2, 1},
		{4, 6},
		{24, 276},
		{25, 276}, // 24 is sum of abundants (12 + 12)
		{26, 301},
		{28123, 4179871},
	}

	for _, ex := range examples {
		result := Run(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
