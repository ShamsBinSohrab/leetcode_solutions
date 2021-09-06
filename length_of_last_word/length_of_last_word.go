package main

import (
	"fmt"
	"strings"
)

//Problem: https://leetcode.com/problems/length-of-last-word/

func main() {
	s := "a"
	fmt.Println(lengthOfLastWord(s))

}

func lengthOfLastWord(s string) int {
	arr := strings.Split(s, " ")
	for i := len(arr) - 1; i >= 0; i-- {
		if len(arr[i]) > 0 {
			return len(arr[i])
		}
	}
	return 0
}
