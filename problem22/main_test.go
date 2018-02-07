package main

import (
	"reflect"
	"testing"
)

var examples = []struct {
	input  string
	answer int
}{
	{`"MARY", "PATRICIA", "LINDA"`, 0},
}

func TestRun(t *testing.T) {
	for _, ex := range examples {
		result := Run(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
