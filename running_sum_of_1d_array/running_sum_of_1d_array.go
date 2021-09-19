package main

import "fmt"

//Problem: https://leetcode.com/problems/running-sum-of-1d-array/

func main() {
	nums := []int{3, 1, 2, 10, 1}
	fmt.Println(runningSum2(nums))
}

func runningSum(nums []int) []int {
	var sum int
	var array []int
	for _, num := range nums {
		sum += num
		array = append(array, sum)
	}
	return array
}

func runningSum2(nums []int) []int {
	var i, j = 0, 1
	for i < len(nums)-1 {
		nums[j] = nums[i] + nums[j]
		i++
		j++
	}
	return nums
}
