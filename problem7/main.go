// Package problem7 By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// What is the 10 001st prime number?
package main

import (
	"fmt"

	"github.com/pwmcintyre/projecteuler/prime"
)

// Run run
func Run(length int) int {
	p := prime.Length(length)
	return p[len(p)-1]
}

func main() {
	p := Run(10001)
	fmt.Println(p)
}
