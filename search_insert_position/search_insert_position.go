package main

import "fmt"

//Problem: https://leetcode.com/problems/search-insert-position/

func main() {
	nums := []int{1, 3}
	target := 0
	fmt.Println(searchInsert(nums, target))
}

var binSearch func([]int, int, int, int) int

func searchInsert(nums []int, target int) int {
	binSearch = func(nums []int, left int, right int, target int) int {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if left == mid && nums[mid] != target {
			if nums[mid] < target {
				return right
			} else {
				return right - 1
			}
		}
		if nums[mid] < target {
			return binSearch(nums, mid, right, target)
		} else {
			return binSearch(nums, left, mid, target)
		}
	}
	return binSearch(nums, 0, len(nums), target)
}
