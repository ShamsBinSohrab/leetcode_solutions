package main

//Problem: https://leetcode.com/problems/remove-duplicates-from-sorted-list/

func main() {
	root := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}
	deleteDuplicates(root)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

var traverse func(*ListNode)

func deleteDuplicates(head *ListNode) *ListNode {
	var prevNode *ListNode
	traverse = func(node *ListNode) {
		if node == nil {
			return
		}
		if prevNode != nil && prevNode.Val == node.Val {
			prevNode.Next = node.Next
		} else {
			prevNode = node
		}
		traverse(node.Next)
	}
	traverse(head)
	return head
}
