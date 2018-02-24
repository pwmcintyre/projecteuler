// package main
/*

By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.

3
7 4
2 4 6
8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom of the triangle below:

75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23

NOTE: As there are only 16384 routes, it is possible to solve this problem by trying every route. However, Problem 67, is the same challenge with a triangle containing one-hundred rows; it cannot be solved by brute force, and requires a clever method! ;o)

Unsure how to approach this, I found a novel solution to optimizing here:
http://www.mathblog.dk/project-euler-18/

*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Run run
func Run(text string) int64 {

	numbers := Parse(text)

	// work from bottom to top
	// but starting from the second bottom
	for x := len(numbers) - 2; x >= 0; x-- {

		// for each number, add the maximum of
		// pair below it
		for y := range numbers[x] {
			numbers[x][y] += Max(numbers[x+1][y], numbers[x+1][y+1])
		}
	}

	// now you've reduced your way to the top
	return numbers[0][0]
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func Parse(text string) [][]int64 {

	text = strings.TrimSpace(text)

	numbers := [][]int64{}

	for x, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)
		numbers = append(numbers, []int64{})
		for _, num := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(num)
			numbers[x] = append(numbers[x], int64(n))
		}
	}

	return numbers
}

func main() {
	fmt.Println(Run(`
		75
		95 64
		17 47 82
		18 35 87 10
		20 04 82 47 65
		19 01 23 75 03 34
		88 02 77 73 07 63 67
		99 65 04 28 06 16 70 92
		41 41 26 56 83 40 80 70 33
		41 48 72 33 47 32 37 16 94 29
		53 71 44 65 25 43 91 52 97 51 14
		70 11 33 28 77 73 17 78 39 68 17 57
		91 71 52 38 17 14 91 43 58 50 27 29 48
		63 66 04 68 89 53 67 30 73 16 69 87 40 31
		04 62 98 27 23 09 70 98 73 93 38 53 60 04 23
		`))
}
