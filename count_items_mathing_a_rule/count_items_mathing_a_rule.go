package main

import "fmt"

//Problem: https://leetcode.com/problems/count-items-matching-a-rule/

func main() {
	items := [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}
	ruleKey := "type"
	ruleValue := "phone"
	fmt.Println(countMatches2(items, ruleKey, ruleValue))
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	props := [3]string{"type", "color", "name"}
	m := make(map[string]int)
	for i := 0; i < len(items); i++ {
		for j := 0; j < len(items[i]); j++ {
			m[props[j]+items[i][j]] += 1
		}
	}
	return m[ruleKey+ruleValue]
}

var props = map[string]int{"type": 0, "color": 1, "name": 2}

func countMatches2(items [][]string, ruleKey string, ruleValue string) int {
	var count int
	for _, item := range items {
		if item[props[ruleKey]] == ruleValue {
			count++
		}
	}
	return count
}
