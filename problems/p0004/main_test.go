package main

import (
	"reflect"
	"testing"
)

var examples = []struct {
	input   int
	answer  int
	numbers []int
}{
	{1, 8, []int{4, 2}},
	{2, 9009, []int{91, 99}},
	{3, 906609, []int{91, 99}},
	// {4, 99000099, []int{91, 99}}, // too slow!
}

func TestRun(t *testing.T) {

	for _, ex := range examples {
		result, err := Run(ex.input)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
