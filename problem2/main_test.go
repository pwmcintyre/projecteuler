package problem2

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {

	// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
	examples := []struct {
		input  int
		answer int
	}{
		{2, 2},
		{3, 2},
		{5, 2},
		{8, 10},
		{13, 10},
		{34, 44},
		{4000000, 4613732},
	}

	for _, ex := range examples {
		result := Run(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
