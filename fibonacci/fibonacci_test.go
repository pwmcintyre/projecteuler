package fibonacci

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

var sequence = []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}

func TestLength(t *testing.T) {

	examples := []struct {
		input  int
		answer []int
	}{
		{10, sequence},
	}

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

func TestGenerator(t *testing.T) {

	next := Generator()

	for i := 0; i < len(sequence); i++ {
		n := next()
		if !reflect.DeepEqual(n, sequence[i]) {
			t.Errorf("\nresult: %+v\nexpected: %d\ni: %+v", n, sequence[i], i)
		}
	}
}

func TestBigGenerator(t *testing.T) {

	next := BigGenerator()

	for i := 0; i < len(sequence); i++ {
		n := next()
		answer := big.NewInt(int64(sequence[i]))
		if !reflect.DeepEqual(n, answer) {
			t.Errorf("\nresult: %+v\nexpected: %d\ni: %+v", n, answer, i)
		}
	}
}
