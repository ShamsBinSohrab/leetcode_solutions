package main

import "fmt"

//Problem: https://leetcode.com/problems/implement-strstr/

func main() {
	haystack := "a"
	needle := "a"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	hl := len(haystack)
	nl := len(needle)
	if hl == 0 && nl == 0 {
		return 0
	}
	for i := 0; (i + nl) <= hl; i++ {
		if haystack[i:(i+nl)] == needle {
			return i
		}
	}
	return -1
}
