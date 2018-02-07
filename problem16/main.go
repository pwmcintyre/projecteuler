// package main
// 2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
// What is the sum of the digits of the number 2^1000?
package main

import (
	"fmt"
	"math/big"
	"strconv"
)

// Run run
func Run(n uint) (sum uint) {

	result := big.NewInt(2)

	// 2^n (or 2 << n-1)
	result = result.Lsh(result, n-1)

	str := result.String()
	sum = 0
	for i := range str {
		e, _ := strconv.Atoi(string(str[i]))
		sum += uint(e)
	}

	return sum
}

func main() {
	fmt.Println(Run(1000))
}
