package functions

import (
	"reflect"
	"testing"
)

func TestCirculations(t *testing.T) {

	var test = []struct {
		input  string
		expect []string
	}{
		{"", []string{""}},
		{"A", []string{"A"}},
		{"AB", []string{"AB", "BA"}},
		{"ABC", []string{"ABC", "CAB", "BCA"}},
	}

	for _, tc := range test {
		result := Circulations(tc.input)
		if !reflect.DeepEqual(tc.expect, result) {
			t.Errorf("\nresult: %s\nexpected: %+v", result, tc.expect)
		}
	}

}
