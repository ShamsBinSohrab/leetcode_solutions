package main

import (
	"fmt"
)

func main() {
	fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {
	if n < 1 {
		return 0
	} else if n <= 2 {
		return n
	}
	steps := []int{1, 2}
	for i := 2; i < n; i++ {
		steps = append(steps, steps[i-1]+steps[i-2])
	}
	return steps[n-1]
}
