package functions

import (
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {

	var test = []struct {
		input  string
		expect []string
	}{
		{"", []string{""}},
		{"A", []string{"A"}},
		{"AB", []string{"AB", "BA"}},
		{"ABC", []string{"ABC", "ACB", "BAC", "BCA", "CBA", "CAB"}},
	}

	for _, tc := range test {
		result := Permutations(tc.input)
		if !reflect.DeepEqual(tc.expect, result) {
			t.Errorf("\nresult: %s\nexpected: %+v", result, tc.expect)
		}
	}

}
