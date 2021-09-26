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
		if _, ok := hash[nums[i]]; !ok {
			hash[nums[i]] = idx
			idx++
		}
	}
	for k, v := range hash {
		if v == 2 {
			return k
		}
	}
	return nums[n-1]
}
