package prime

import (
	"reflect"
	"testing"
)

// first 46 primes
var testPrimes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}

func TestLength(t *testing.T) {

	examples := []struct {
		input  int
		answer []int
	}{
		{1, testPrimes[0:1]},
		{2, testPrimes[0:2]},
		{10, testPrimes[0:10]},
		{20, testPrimes[0:20]},
		{46, testPrimes},
	}

	for _, ex := range examples {
		result := Length(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

func benchmarkLength(len int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Length(len)
	}
}
func BenchmarkLength100(b *testing.B)  { benchmarkLength(100, b) }
func BenchmarkLength1000(b *testing.B) { benchmarkLength(1000, b) }

func TestUntil(t *testing.T) {

	examples := []struct {
		input  int
		answer []int
	}{
		{1, testPrimes[0:0]},
		{2, testPrimes[0:1]},
		{10, testPrimes[0:4]},
		{20, testPrimes[0:8]},
		{200, testPrimes},
	}

	for _, ex := range examples {
		result := Until(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
		t.Log(result)
	}
}

func TestSieveOfEratosthenes(t *testing.T) {

	examples := []struct {
		input  int
		answer []int
	}{
		{1, testPrimes[0:0]},
		{2, testPrimes[0:1]},
		{10, testPrimes[0:4]},
		{20, testPrimes[0:8]},
		{200, testPrimes},
	}

	for _, ex := range examples {
		result := SieveOfEratosthenes(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\ninput: %+v", result, ex.answer, ex.input)
		}
	}
}

func BenchmarkSieveOfEratosthenes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SieveOfEratosthenes(n)
	}
}

func TestGenerator(t *testing.T) {

	gen := Generator()

	for i, ex := range testPrimes {
		result := gen()
		if !reflect.DeepEqual(ex, result) {
			t.Errorf("\nresult: %+v\nexpected: %d\nindex: %+v", result, ex, i)
		}
		t.Log(result)
	}
}

func benchmarkGenerator(len int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		gen := Generator()
		for i := 0; i < len; i++ {
			gen()
		}
	}
}

func BenchmarkGenerator1(b *testing.B)    { benchmarkGenerator(1, b) }
func BenchmarkGenerator10(b *testing.B)   { benchmarkGenerator(10, b) }
func BenchmarkGenerator100(b *testing.B)  { benchmarkGenerator(100, b) }
func BenchmarkGenerator1000(b *testing.B) { benchmarkGenerator(1000, b) }

func TestIsPrime(t *testing.T) {

	examples := []struct {
		input  int
		answer bool
	}{
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{199, true},
		{1000, false},
	}

	for _, ex := range examples {
		result := IsPrime(ex.input)
		if !reflect.DeepEqual(ex.answer, result) {
			t.Errorf("\nresult: %+v\nexpected: %v\ninput: %+v", result, ex.answer, ex.input)
		}
		t.Log(result)
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for n := 0; n < b.N; n++ {
		purgeCache()
		for i := 0; i < len(testPrimes); i++ {
			IsPrime(i)
		}
	}
}
