package main

import (
	"fmt"
)

//Problem: https://leetcode.com/problems/linked-list-cycle/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	_3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	_2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	_0 := &ListNode{
		Val:  0,
		Next: nil,
	}
	_4 := &ListNode{
		Val:  -4,
		Next: nil,
	}
	_3.Next = _2
	_2.Next = _0
	_0.Next = _4
	_4.Next = _2

	fmt.Println(hasCycle2(_3))
}

//Brute force
func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]struct{})
	curr := head
	for curr != nil {
		if _, ok := m[curr]; ok {
			return ok
		}
		m[curr] = struct{}{}
		curr = curr.Next
	}
	return false
}

//Fast and Slow pointer
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
