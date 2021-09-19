package main

import (
	"fmt"
	"strings"
)

//Problems: https://leetcode.com/problems/valid-palindrome/

func main() {
	s := "0P"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	var arr []string
	for _, r := range strings.ToLower(s) {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			arr = append(arr, string(r))
		}
	}
	var low, high = 0, len(arr) - 1
	for low < high {
		if arr[low] != arr[high] {
			return false
		}
		low++
		high--
	}
	return true
}
