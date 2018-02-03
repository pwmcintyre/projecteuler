package fibonacci

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLength(t *testing.T) {

	examples := []struct {
		input  int
		answer []int
	}{
		{10, []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}},
	}

	fmt.Print(examples)

	for _, ex := range examples {
		result := Length(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

func TestUntil(t *testing.T) {

	examples := []struct {
		input  int
		answer []int
	}{
		{5, []int{1, 2, 3, 5}},
		{90, []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}},
	}

	fmt.Print(examples)

	for _, ex := range examples {
		result := Until(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}
