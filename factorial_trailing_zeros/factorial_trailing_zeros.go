package main

import "fmt"

//Problem: https://leetcode.com/problems/factorial-trailing-zeroes/

func main() {
	fmt.Println(trailingZeroes(100))
}

func trailingZeroes(n int) int {
	count := 0
	for n > 0 {
		count += n / 5
		n /= 5
	}
	return count
}
