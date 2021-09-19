package main

import "fmt"

//Problem: https://leetcode.com/problems/shuffle-the-array/

func main() {
	nums := []int{2, 5, 1, 3, 4, 7}
	n := 3
	fmt.Println(shuffle(nums, n))
}

func shuffle(nums []int, n int) []int {
	array := make([]int, len(nums))
	for i := 0; i < n; i++ {
		array[2*i] = nums[i]
		array[2*i+1] = nums[i+n]
	}
	return array
}
