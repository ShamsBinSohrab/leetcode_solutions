package main

import "fmt"

//Problem: https://leetcode.com/problems/shuffle-the-array/

func main() {
	nums := []int{1, 2, 3, 4, 4, 3, 2, 1}
	n := 4
	fmt.Println(shuffle(nums, n))
}

//In place
func shuffle(nums []int, n int) []int {
	for i := 0; i < n; i++ {
		insertAt(2*i+1, nums[2*i+n], &nums)
	}
	return nums[:len(nums)-n]
}

func insertAt(index, element int, ptr *[]int) {
	*ptr = append(*ptr, 0)
	copy((*ptr)[index+1:], (*ptr)[index:])
	(*ptr)[index] = element
}

//Another array easy solution
func shuffle2(nums []int, n int) []int {
	array := make([]int, len(nums))
	for i := 0; i < n; i++ {
		array[2*i] = nums[i]
		array[2*i+1] = nums[i+n]
	}
	return array
}
