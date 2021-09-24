package main

import "fmt"

//Problem: https://leetcode.com/problems/excel-sheet-column-number/

func main() {
	columnTitle := "FXSHRXW"
	fmt.Println(titleToNumber(columnTitle))
}

func titleToNumber(columnTitle string) int {
	number := 0
	for _, letter := range columnTitle {
		number = (number * 26) + int(letter) - 64
	}
	return number
}
