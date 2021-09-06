package main

import "sort"

//Problem: https://leetcode.com/problems/merge-two-sorted-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := ListNode{}
	l2 := ListNode{}

	mergeTwoLists(&l1, &l2)
}

var traverse func(*ListNode)

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	vals := make([]int, 0)
	traverse = func(node *ListNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		traverse(node.Next)
	}
	traverse(l1)
	traverse(l2)
	sort.Sort(sort.Reverse(sort.IntSlice(vals)))

	var root = &ListNode{}
	var node *ListNode
	var prevNode *ListNode

	for i, v := range vals {
		node = &ListNode{
			Val:  v,
			Next: prevNode,
		}
		prevNode = node
		if i == len(vals) -1 {
			root = node
		}
	}
	return root
}
