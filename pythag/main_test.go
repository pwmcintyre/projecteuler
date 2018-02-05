package pythag

import (
	"reflect"
	"testing"
)

var examples = []struct {
	input  int
	answer [][]int
}{
	{4, [][]int{}},
	{5, [][]int{{3, 4, 5}}},
	{6, [][]int{{3, 4, 5}}},
	{10, [][]int{{3, 4, 5}, {6, 8, 10}}},
}

func TestPythagoreanTriplets(t *testing.T) {
	for _, ex := range examples {
		result := Until(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
