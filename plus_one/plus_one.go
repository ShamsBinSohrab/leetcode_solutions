package main

import "fmt"

//Problem: https://leetcode.com/problems/plus-one/

func main() {
	digits := []int{0}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	var num int
	var carry = 1
	for i := len(digits) - 1; i >= 0; i-- {
		num = digits[i] + carry
		digits[i] = num % 10
		carry = num / 10
		if i == 0 && carry > 0 {
			digits = append([]int{carry}, digits...)
		}
	}
	return digits
}
