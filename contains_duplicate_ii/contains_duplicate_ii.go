package main

import "fmt"

//Problems: https://leetcode.com/problems/contains-duplicate-ii/

func main() {
	nums := []int{1, 0, 1, 1}
	k := 1
	fmt.Println(containsNearbyDuplicate(nums, k))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if v, ok := m[num]; ok && abs(v-i) <= k {
			return true
		} else {
			m[num] = i
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
