package main

import (
	"fmt"
	"sort"
)

//Problem: https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/

func main() {
	nums := []int{7, 7, 7, 7}
	fmt.Println(smallerNumbersThanCurrent(nums))
}

func smallerNumbersThanCurrent(nums []int) []int {
	var result = make([]int, len(nums))
	var temp = make([]int, len(nums))
	copy(temp, nums)
	sort.Ints(temp)
	for i, num := range nums {
		result[i] = indexOf(num, temp)
	}
	return result
}

func indexOf(n int, arr []int) int {
	for i, el := range arr {
		if el == n {
			return i
		}
	}
	return -1
}
