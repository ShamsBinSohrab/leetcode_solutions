package main

import "fmt"

//Problem: https://leetcode.com/problems/excel-sheet-column-title/

func main() {
	columnNumber := 28
	fmt.Println(convertToTitle(columnNumber))
}

func convertToTitle(columnNumber int) string {
	result := ""
	for columnNumber > 0 {
		columnNumber--
		result = string((columnNumber%26)+65) + result
		columnNumber /= 26
	}
	return result
}
