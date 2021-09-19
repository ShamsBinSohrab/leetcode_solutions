package main

import "fmt"

//Problem: https://leetcode.com/problems/shuffle-string/

func main() {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	fmt.Println(restoreString(s, indices))
}

func restoreString(s string, indices []int) string {
	var arr = make([]uint8, len(s))
	for i, index := range indices {
		arr[index] = s[i]
	}
	var result string
	for _, ar := range arr {
		result += string(ar)
	}
	return result
}
