package main

import (
	"fmt"
)

func Sum(x int, y int) int {
	return x + y
}

func main() {
	result := Sum(5, 5)
	fmt.Println(result)
}
