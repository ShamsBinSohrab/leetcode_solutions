package main

import "fmt"

//Problem: https://leetcode.com/problems/number-of-good-pairs/

func main() {
	nums := []int{1, 1, 1, 1}
	//nums := []int{1,2,3,1,1,3}
	fmt.Println(numIdenticalPairs2(nums))

}

//O(n^2)
func numIdenticalPairs(nums []int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
	}
	return count
}

//O(n)
func numIdenticalPairs2(nums []int) int {
	var count int
	var m = make(map[int]int)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			count += m[num]
		}
		m[num]++
	}
	return count
}
