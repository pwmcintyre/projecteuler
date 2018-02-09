package fibonacci

import (
	"math/big"
	"reflect"
	"testing"
)

var sequence = []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}

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
