package main

import (
	"fmt"
)

//Problem: https://leetcode.com/problems/valid-parentheses/

func main() {
	s := "]"
	fmt.Println(isValid(s))
}

var m = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
}

type stack struct {
	list []string
}

func (stk *stack) push(s string) {
	stk.list = append(stk.list, s)
}

func (stk *stack) pop() string {
	top := len(stk.list) - 1
	if top < 0 {
		return ""
	}
	elem := stk.list[top]
	stk.list = stk.list[:top]
	return elem
}

func (stk *stack) isEmpty() bool {
	return len(stk.list) == 0
}

func isValid(s string) bool {
	stk := stack{list: make([]string, 0)}
	for _, v := range s {
		p := string(v)
		if _, ok := m[p]; ok {
			stk.push(p)
		} else if m[stk.pop()] != p {
			return false
		}
	}
	return stk.isEmpty()
}
