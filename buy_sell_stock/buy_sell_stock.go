package main

import (
	"fmt"
)

//Problems: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	var i, buy, max = 1, prices[0], 0
	for i < len(prices) {
		profit := prices[i] - buy
		max = getMax(profit, max)
		buy = getMin(prices[i], buy)
		i++
	}
	return max
}

func getMax(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func getMin(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
