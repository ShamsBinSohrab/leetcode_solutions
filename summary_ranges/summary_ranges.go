package main

import (
	"fmt"
	"strconv"
)

//Problem: https://leetcode.com/problems/summary-ranges/

func main() {
	nums := []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println(summaryRanges(nums))
}

func summaryRanges(nums []int) []string {
	var result []string
	size := len(nums)
	if size > 0 {
		ptr := nums[0]
		for i := 0; i < size; i++ {
			if i == size-1 || nums[i]+1 != nums[i+1] {
				if nums[i] != ptr {
					result = append(result, fmt.Sprintf("%s->%s", strconv.Itoa(ptr), strconv.Itoa(nums[i])))
				} else {
					result = append(result, strconv.Itoa(ptr))
				}
				if i != size-1 {
					ptr = nums[i+1]
				}
			}
		}
	}
	return result
}
