package main

import (
	"fmt"
	"sort"
)

//Problem: https://leetcode.com/problems/sign-of-the-product-of-an-array/

func main() {
	nums := []int{-1, 1, -1, 1, -1}
	fmt.Println(arraySign(nums))
}

func arraySign(nums []int) int {
	sort.Ints(nums)
	flag := 1
	for _, v := range nums {
		if v < 0 {
			flag = -flag
		} else if v == 0 {
			return 0
		}
	}
	return flag
}
