package problem5

import (
	"reflect"
	"testing"
)

var examples = []struct {
	input  Input
	answer int
}{
	{Input{1, 2}, 2},
	{Input{1, 3}, 6},
	{Input{1, 4}, 12},
	{Input{1, 5}, 60},
	{Input{1, 10}, 2520},
	{Input{1, 20}, 232792560},
}

func TestRun(t *testing.T) {

	for _, ex := range examples {
		result, err := Run(ex.input)
		if err != nil {
			t.Error(ex.input, err)
		} else if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
