package main

import (
	"fmt"
)

//Problem: https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

func main() {
	nums := []int{1, 2, 5, 5, 6, 6}
	fmt.Println(findDisappearedNumbers(nums))
}

func findDisappearedNumbers(nums []int) []int {
	for _, v := range nums {
		if v < 0 {
			v = -1 * v
		}
		if nums[v-1] > 0 {
			nums[v-1] = -1 * nums[v-1]
		}
	}
	var res []int
	for i, v := range nums {
		if v > 0 {
			res = append(res, i+1)
		}
	}
	return res
}
