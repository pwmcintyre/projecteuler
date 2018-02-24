package main

import (
	"reflect"
	"testing"
)

var examples = []struct {
	input  uint
	answer uint
}{
	{1, 2},
	{2, 4},
	{3, 8},
	{4, 7},       // 16 > 1 + 6 = 7
	{15, 26},     // 32768 > 3 + 2 + 7 + 6 + 8 = 26
	{1000, 1366}, // 10715086071862673209484250490600018105614048117055336074437503883703510511249361224931983788156958581275946729175531468251871452856923140435984577574698574803934567774824230985421074605062371141877954182153046474983581941267398767559165543946077062914571196477686542167660429831652624386837205668069376 > ...
}

func TestRun(t *testing.T) {
	for _, ex := range examples {
		result := Run(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
