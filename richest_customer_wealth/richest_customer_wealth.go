package main

import "fmt"

//Problem: https://leetcode.com/problems/richest-customer-wealth/
func main() {
	accounts := [][]int{{1, 2, 3}, {3, 2, 1}}
	fmt.Println(maximumWealth(accounts))
}

func maximumWealth(accounts [][]int) int {
	var l1, l2, max int
	array := make([]int, len(accounts))
	for l1 < len(accounts[0]) {
		for l2 < len(accounts) {
			array[l2] += accounts[l2][l1]
			max = getMax(max, array[l2])
			l2++
		}
		l1++
		l2 = 0
	}
	return max
}

func getMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
