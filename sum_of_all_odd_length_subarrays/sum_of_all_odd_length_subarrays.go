package main

import "fmt"

//Problem: https://leetcode.com/problems/sum-of-all-odd-length-subarrays/
//Explanation: https://www.youtube.com/watch?v=J5IIH35EBVE&t=616s

func main() {
	arr := []int{1, 4, 2, 5, 3}
	fmt.Println(sumOddLengthSubarrays(arr))
}

func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	sum := 0
	for i := 0; i < n; i++ {
		startWith := n - i
		endWith := i + 1
		if total := startWith * endWith; total%2 == 1 {
			sum += (total/2 + 1) * arr[i]
		} else {
			sum += total / 2 * arr[i]
		}
	}
	return sum
}
func sumOddLengthSubarrays2(arr []int) int {
	sum := 0
	l := len(arr)
	for a := 0; a < l; a++ {
		for b := a; b < l; b += 2 {
			for c := a; c <= b; c++ {
				sum += arr[c]
			}
		}
	}
	return sum
}
