package main

import "fmt"

//Problem: https://leetcode.com/problems/check-if-array-is-sorted-and-rotated/

func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(check(nums))
}

func check(nums []int) bool {
	count := 0
	for i, v := range nums {
		if v > nums[(i+1)%len(nums)] {
			count += 1
			if count > 1 {
				return false
			}
		}
	}
	return true
}
