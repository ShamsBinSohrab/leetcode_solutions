package main

import "fmt"

//Problems:

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumberXor(nums))
}

func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num] += 1
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

func singleNumberXor(nums []int) int {
	res := 0
	for _, num := range nums {
		res = res ^ num
	}
	return res
}
