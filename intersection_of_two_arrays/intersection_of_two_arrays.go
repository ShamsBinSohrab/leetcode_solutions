package main

import (
	"fmt"
)

//Problem: https://leetcode.com/problems/intersection-of-two-arrays/

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {
	hash := make(map[int]bool)
	var common []int
	for _, i := range nums1 {
		hash[i] = false
	}
	for _, i := range nums2 {
		if inserted, ok := hash[i]; ok && !inserted {
			hash[i] = true
			common = append(common, i)
		}
	}
	return common
}
