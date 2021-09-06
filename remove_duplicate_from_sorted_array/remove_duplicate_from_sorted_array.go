package main

import (
	"fmt"
)

//Problem: https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			nums = append(nums[:i-1], nums[i:]...)
			i--
		}
	}
	return len(nums)
}
