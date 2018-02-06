package factorial

import (
	"math/big"
	"reflect"
	"strconv"
	"testing"
)

// http://2000clicks.com/mathhelp/BasicFactorialTable.aspx
var examples = []struct {
	input  uint
	answer string
}{
	{1, "1"},
	{2, "2"},
	{3, "6"},
	{4, "24"},
	{10, "3628800"},
	{20, "2432902008176640000"},
	{30, "265252859812191058636308480000000"},
	{50, "30414093201713378043612608166064768844377641568960512000000000000"},
	{100, "93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000"},
}

func TestFactorial(t *testing.T) {
	for _, ex := range examples[0:6] {

		a, err := strconv.Atoi(ex.answer)
		if err != nil {
			t.Error(err)
		}
		answer := uint(a)
		result := Factorial(ex.input)

		if !reflect.DeepEqual(answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, answer, ex.input)
		}
	}
}

func TestBigFactorial(t *testing.T) {
	for _, ex := range examples {

		i := big.NewInt(int64(ex.input))

		a := new(big.Int)
		a.SetString(ex.answer, 10)

		result := BigFactorial(i)

		if !reflect.DeepEqual(a, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, a, i)
		}
	}
}
