package main

import "fmt"

//Problem: https://leetcode.com/problems/move-zeroes/

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	next := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[next], nums[i] = nums[i], nums[next]
			next++
		}
	}
}
