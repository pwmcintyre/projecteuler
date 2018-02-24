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
	{3, 3},
	{6, 7},
	{50, 224},
	{200, 2015},
	{300, 2079},
	// {400, 5984},
	// {500, 12375}, // too slow
}

func TestRun(t *testing.T) {
	for _, ex := range examples {
		result, _ := Run(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
