// package main
/*

A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions with denominators 2 to 10 are given:

1/2	= 	0.5
1/3	= 	0.(3)
1/4	= 	0.25
1/5	= 	0.2
1/6	= 	0.1(6)
1/7	= 	0.(142857)
1/8	= 	0.125
1/9	= 	0.(1)
1/10	= 	0.1
Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.

Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.

*/

package main

import (
	"fmt"
	"strconv"

	"github.com/pwmcintyre/projecteuler/arrays"
)

// Run run
func Run(below int) (answer int) {

	max := 0
	for i := below; i > 0 && i > max; i-- {

		_, _, cycle := GetCycle(1, i, i)
		if len(cycle) > max {
			max = len(cycle)
			answer = i
		}

		fmt.Printf("i: %d cycle: %d max: %d maxi: %d\n", i, len(cycle), max, answer)
	}

	return
}

// GetCycle takes a fraction as two integers, numerator and denominator, and returns details about recuring cycles.
// It essentially performs long division, keeping track of digits it has already divided.
// https://en.wikipedia.org/wiki/Repeating_decimal
// http://virtualnerd.com/middle-math/number-theory-fractions/fractions-decimals/fraction-to-repeating-decimal-conversion
func GetCycle(numerator, denominator int, precision int) (start int, str, cycle string) {

	//
	remainder := []int{numerator}
	dividends := []int{}

	done := false

	for i := 0; !done && numerator != 0 && i < precision; i++ {

		numerator *= 10
		dividends = append(dividends, int(numerator/denominator))
		numerator %= denominator

		done, start = arrays.InArray(numerator, remainder)
		remainder = append(remainder, numerator)

	}

	if done {
		for _, n := range dividends[start:] {
			cycle += strconv.Itoa(n)
		}
	}

	return
}

func main() {
	fmt.Println(Run(1000))
}
