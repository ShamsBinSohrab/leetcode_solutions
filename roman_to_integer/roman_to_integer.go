package main

import "fmt"

//Problem: https://leetcode.com/problems/roman-to-integer/

func main() {
	roman := "MCMXCIV"
	fmt.Println(romanToInt(roman))
}

var m = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	sum := 0
	skipNext := false
	for i, v := range s {
		if skipNext {
			skipNext = false
			continue
		}
		l := string(v)
		v1 := m[l]
		if i != (len(s) - 1) {
			nl := string(s[i+1])
			v2 := m[nl]
			if v1 >= v2 {
				sum += v1
			} else {
				skipNext = true
				sum += v2 - v1
			}
		} else {
			sum += v1
		}
	}
	return sum
}
