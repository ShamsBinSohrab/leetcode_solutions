package main

import "fmt"

//Problem: https://leetcode.com/problems/build-array-from-permutation/

func main() {
	nums := []int{0, 2, 1, 5, 3, 4}
	fmt.Println(buildArray(nums))
}

func buildArray(nums []int) []int {
	var array []int
	for _, num := range nums {
		array = append(array, nums[num])
	}
	return array
}
