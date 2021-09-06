package main

import (
	"fmt"
	"strconv"
)

//Problem: https://leetcode.com/problems/add-binary/

func main() {
	a := "11"
	b := "1"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	if diff := len(a) - len(b); diff > 0 {
		for diff != 0 {
			b = "0" + b
			diff--
		}
	} else if diff < 0 {
		for diff != 0 {
			a = "0" + a
			diff++
		}
	}
	var sum, carry int
	var result string
	var c = len(a) - 1
	for c >= 0 {
		sum = int(a[c]-'0') + int(b[c]-'0') + carry
		result = strconv.Itoa(sum%2) + result
		carry = sum / 2
		c--
	}
	if carry != 0 {
		result = strconv.Itoa(carry) + result
	}
	return result
}
