package main

import (
	"fmt"
	"sort"
)

//Problem: https://leetcode.com/problems/assign-cookies/

func main() {
	g := []int{7, 8, 9, 10}
	s := []int{5, 6, 7, 8}
	fmt.Println(findContentChildren(g, s))
}

func findContentChildren(g []int, s []int) int {
	gptr, sptr := 0, 0
	sort.Ints(g)
	sort.Ints(s)
	for gptr < len(g) && sptr < len(s) {
		if g[gptr] <= s[sptr] {
			gptr++
			sptr++
		} else {
			sptr++
		}
	}
	return gptr
}
