package main

import (
	"fmt"
	"sort"
)

//Problems: https://leetcode.com/problems/contains-duplicate/

func main() {
	nums := []int{0, 4, 5, 0, 3, 6}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
