package main

import (
	"fmt"
	"sort"
)

//Problem: https://leetcode.com/problems/third-maximum-number/

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(thirdMax(nums))
}

func thirdMax(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	if n < 3 {
		return nums[n-1]
	}
	idx := 0
	hash := make(map[int]int)
	for i := n - 1; i >= 0 && idx < 3; i-- {
		if idx == 0 || hash[idx-1] != nums[i] {
			hash[idx] = nums[i]
			idx++
		}
	}
	if v, ok := hash[2]; ok {
		return v
	}
	return nums[n-1]
}
