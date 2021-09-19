package main

import (
	"fmt"
)

//Problem: https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/

func main() {
	prices := []int{10, 1, 1, 6}
	fmt.Println(finalPrices(prices))
}

func finalPrices(prices []int) []int {
	final := make([]int, 0, len(prices))
	for i, price := range prices {
		final = append(final, price)
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= price {
				final[i] = price - prices[j]
				break
			}
		}
	}
	return final
}
