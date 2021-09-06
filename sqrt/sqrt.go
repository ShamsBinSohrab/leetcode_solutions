package main

import "fmt"

//Problem: https://leetcode.com/problems/sqrtx/

func main() {
	c := 2147483647
	fmt.Println(mySqrt(c))
}

var binSearchSqrt func(int, int, int) int

func mySqrt(x int) int {
	var div int
	var mid int
	binSearchSqrt = func(start int, end int, target int) int {
		if target == 1 || target == 2 {
			return 1
		}
		mid = (start + end) / 2
		if mid == start {
			return mid
		} else {
			div = target / mid
			if div%mid <= 1 {
				return mid
			}
		}
		return binSearchSqrt(div, mid, target)
	}
	return binSearchSqrt(0, x, x)
}
