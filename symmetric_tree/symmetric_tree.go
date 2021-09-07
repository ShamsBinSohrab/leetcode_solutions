package main

import "fmt"

//Problem: https://leetcode.com/problems/symmetric-tree/

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(isSymmetric(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var isSameTree func(*TreeNode, *TreeNode) bool

func isSymmetric(root *TreeNode) bool {
	isSameTree = func(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p != nil && q != nil {
			return p.Val == q.Val && isSameTree(p.Left, q.Right) && isSameTree(p.Right, q.Left)
		}
		return false
	}
	return isSameTree(root.Left, root.Right)
}
