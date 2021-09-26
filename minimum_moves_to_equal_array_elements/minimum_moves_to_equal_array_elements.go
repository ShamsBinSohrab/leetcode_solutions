package main

import (
	"fmt"
	"sort"
)

//Problem: https://leetcode.com/problems/minimum-moves-to-equal-array-elements/

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(minMoves(nums))
}

func minMoves(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	sort.Ints(nums)
	moves := 0
	for _, num := range nums {
		moves += num - nums[0]
	}
	return moves
}
