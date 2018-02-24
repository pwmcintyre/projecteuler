package main

import (
	"reflect"
	"testing"
)

var examples = []struct {
	start  int
	end    int
	answer int
}{
	{1901, 2001, 171},
}

func TestRun(t *testing.T) {
	for _, ex := range examples {
		result := Run(ex.start, ex.end)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.start)
		}
	}
}
