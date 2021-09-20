package main

import (
	"fmt"
)

//Problem: https://leetcode.com/problems/intersection-of-two-linked-lists/

func main() {
	a1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	a9 := &ListNode{
		Val:  9,
		Next: nil,
	}
	_a1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	b3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	_2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	_4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	a1.Next = a9
	a9.Next = _a1
	_a1.Next = _2
	b3.Next = _2
	_2.Next = _4

	fmt.Println(getIntersectionNode2(a1, b3))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	for headA != nil {
		m[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := m[headB]; ok {
			return headB
		}
		m[headB] = struct{}{}
		headB = headB.Next
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	countA := countNodes(headA)
	countB := countNodes(headB)
	if countA > countB {
		return findIntersection(headA, headB, countB, countA)
	}
	return findIntersection(headB, headA, countA, countB)
}

func findIntersection(nodeA, nodeB *ListNode, countA, countB int) *ListNode {
	for countA != countB {
		countA++
		nodeA = nodeA.Next
	}
	for nodeA != nil && nodeB != nil {
		if nodeA == nodeB {
			return nodeA
		}
		nodeA = nodeA.Next
		nodeB = nodeB.Next
	}
	return nil
}

func countNodes(node *ListNode) int {
	var count int
	for node != nil {
		count++
		node = node.Next
	}
	return count
}
