// package main
/*

You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

*/

package main

import (
	"fmt"
	"time"
)

// Run run
func Run(a, b int) int {

	utc, _ := time.LoadLocation("UTC")
	start := time.Date(a, 1, 1, 0, 0, 0, 0, utc)
	end := time.Date(b, 1, 1, 0, 0, 0, 0, utc)

	count := 0
	for start.Before(end) {
		start = start.AddDate(0, 0, 1)
		if start.Weekday() == time.Sunday && start.Day() == 1 {
			fmt.Println(start.Format("2006 Jan 01 02 - Monday"))
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(Run(1901, 2001))
}
