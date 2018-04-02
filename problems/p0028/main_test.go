package main

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {

	var examples = []struct {
		length int
		answer int
	}{
		{1, 1},
		{3, 25},
		{5, 101},
	}

	for _, ex := range examples {
		result := Run(ex.length)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\nlimit: %+v", result, ex.answer, ex.limit)
		}
	}
}
