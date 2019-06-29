package functions

import (
	"reflect"
	"testing"
)

func TestTruncationsOf(t *testing.T) {

	var examples = []struct {
		n           int
		truncations []int
	}{
		{1, []int{1}},
		{10, []int{10, 0, 1}},
		{11, []int{11, 1}},
		{123, []int{123, 23, 3, 1, 12}},
	}

	for _, ex := range examples {
		result := TruncationsOf(ex.n)
		if !reflect.DeepEqual(ex.truncations, result) {
			t.Errorf("\nresult: %+v\nexpected: %v\nn: %+v", result, ex.truncations, ex.n)
		}
	}
}
