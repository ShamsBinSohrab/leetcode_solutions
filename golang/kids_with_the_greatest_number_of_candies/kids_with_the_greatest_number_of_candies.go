package main

import "fmt"

//Problem: https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	fmt.Println(kidsWithCandies(candies, extraCandies))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var max int
	for _, candy := range candies {
		max = getMax(max, candy)
	}

	var result = make([]bool, len(candies))
	for i, candy := range candies {
		result[i] = candy+extraCandies >= max
	}
	return result
}

func getMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
