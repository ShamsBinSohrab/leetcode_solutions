package main

import "fmt"

//Problems: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 26
	fmt.Println(twoSum(numbers, target))
}

func twoSum(numbers []int, target int) []int {
	var low, high = 0, len(numbers) - 1
	for low < high {
		if sum := numbers[low] + numbers[high]; sum == target {
			return []int{low + 1, high + 1}
		} else if sum < target {
			low++
		} else {
			high--
		}
	}
	return nil
}
