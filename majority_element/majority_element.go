package main

import "fmt"

//Problems: https://leetcode.com/problems/majority-element/

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}

//Moore's voting algorithm
func majorityElement(nums []int) int {
	var count, candidate = 0, -1
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
			count = 1
		} else if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}
	return candidate
}
