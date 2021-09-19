package main

import "fmt"

//Problem: https://leetcode.com/problems/merge-two-sorted-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	ans := mergeTwoLists(l1, l2)
	fmt.Println(ans)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		if l2 == nil {
			return nil
		}
		return &ListNode{
			Val:  l2.Val,
			Next: mergeTwoLists(nil, l2.Next),
		}
	} else if l2 == nil {
		if l1 == nil {
			return nil
		}
		return &ListNode{
			Val:  l1.Val,
			Next: mergeTwoLists(l1.Next, nil),
		}
	} else if l1.Val > l2.Val {
		return &ListNode{
			Val:  l2.Val,
			Next: mergeTwoLists(l1, l2.Next),
		}
	}
	return &ListNode{
		Val:  l1.Val,
		Next: mergeTwoLists(l1.Next, l2),
	}
}
