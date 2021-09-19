package main

import (
	"fmt"
	"math"
)

//Problem: https://leetcode.com/problems/reverse-integer/

func main() {
	x := -123
	fmt.Println(reverse(x))
}
func reverse(x int) int {
	var rev, mod int
	for x != 0 {
		mod = x % 10
		rev = (rev * 10) + mod
		if rev > math.MaxInt32 || rev < math.MinInt32 {
			return 0
		}
		x /= 10
	}
	return rev
}
