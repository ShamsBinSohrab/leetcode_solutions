package main

import "fmt"

//Problem: https://leetcode.com/problems/concatenation-of-array/

func main() {
	nums := []int{1, 3, 2, 1}
	fmt.Println(getConcatenation(nums))
}

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}

//Solution with an array
func getConcatenationArray(nums [4]int) [8]int {
	array := [len(nums) * 2]int{}
	for i, num := range nums {
		array[i] = num
		array[i+len(nums)] = num
	}
	return array
}
