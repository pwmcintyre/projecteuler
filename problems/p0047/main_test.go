package main

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {

	var examples = []struct {
		input   int
		factors []int
		answer  int
	}{
		{2, []int{14, 15}, 14},
		{3, []int{644, 645, 646}, 644},
		{4, []int{134043, 134044, 134045, 134046}, 134043},
	}

	for _, ex := range examples {
		result := Run(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
