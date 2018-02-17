package fibonacci

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

var (
	sequence = []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	fib92    = 7540113804746346429
	fib1000  = "43466557686937456435688527675040625802564660517371780402481729089536555417949051890403879840079255169295922593080322634775209689623239873322471161642996440906533187938298969649928516003704476137795166849228875"
)

func TestFib(t *testing.T) {

	next := Fib()

	for i := 0; i < len(sequence); i++ {
		n, err := next()
		if err != nil {
			t.Error("unexpected error")
		}
		if n != sequence[i] {
			t.Errorf("\nresult: %+v\nexpected: %d\ni: %+v", n, sequence[i], i)
		}
	}
}

func TestFib92(t *testing.T) {

	next := Fib()

	for i := 1; i < 92; i++ {
		next()
	}

	// maximum
	n, err := next()
	if n != fib92 {
		t.Errorf("\nresult: %d\nexpected: %d", n, fib92)
	}
	if err != nil {
		t.Error("expected an overflow error")
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		next := Fib()
		for j := 0; j < 92; j++ {
			next()
		}
	}
}

func TestBigFib(t *testing.T) {

	next := BigFib()

	for i := 0; i < len(sequence); i++ {
		n := next()
		answer := big.NewInt(int64(sequence[i]))
		if !reflect.DeepEqual(n, answer) {
			t.Errorf("\nresult: %+v\nexpected: %d\ni: %+v", n, answer, i)
		}
	}
}

func BenchmarkBigFib(b *testing.B) {
	for _, i := range []int{1, 100, 10000} {
		b.Run(fmt.Sprintf("Fib-%d", i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				next := BigFib()
				for j := 0; j < i; j++ {
					next()
				}
			}
		})
	}
}
