package main

import "fmt"

//Problem: https://leetcode.com/problems/create-target-array-in-the-given-order/

func main() {
	nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}
	fmt.Println(createTargetArray(nums, index))
}

func createTargetArray(nums []int, index []int) []int {
	var output []int
	for i := 0; i < len(nums); i++ {
		insertAt(index[i], nums[i], &output)
	}
	return output
}

func insertAt(index, element int, ptr *[]int) {
	*ptr = append(*ptr, 0)
	copy((*ptr)[index+1:], (*ptr)[index:])
	(*ptr)[index] = element
}
