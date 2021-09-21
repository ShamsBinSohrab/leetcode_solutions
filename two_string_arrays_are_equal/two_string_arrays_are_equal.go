package main

import (
	"fmt"
	"strings"
)

//Problem: https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/

func main() {
	word1 := []string{"ab", "c"}
	word2 := []string{"a", "bc"}
	fmt.Println(arrayStringsAreEqual(word1, word2))
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	sb1, sb2 := strings.Builder{}, strings.Builder{}
	for _, s := range word1 {
		sb1.WriteString(s)
	}
	for _, s := range word2 {
		sb2.WriteString(s)
	}
	return sb1.String() == sb2.String()
}
