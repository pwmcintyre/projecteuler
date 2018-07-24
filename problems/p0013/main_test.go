package main

import (
	"reflect"
	"testing"
)

var examples = []struct {
	input  int
	answer string
}{
	{1, "5537376230390876637302048746832985971773659831892672"},
}

func TestRun(t *testing.T) {
	for _, ex := range examples {
		result := Run().String()
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %v\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
