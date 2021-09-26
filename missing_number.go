package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 0, 1}
	fmt.Println(missingNumberXOR(nums))
}

func missingNumber(nums []int) int {
	size := len(nums)
	sum := size * (size + 1) / 2
	for i := 0; i < size; i++ {
		sum -= nums[i]
	}
	return sum
}

func missingNumberXOR(nums []int) int {
	var result int
	for i, num := range nums {
		result ^= i ^ num
	}
	return result ^ len(nums)
}
