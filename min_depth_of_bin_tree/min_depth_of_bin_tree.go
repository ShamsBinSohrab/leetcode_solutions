package main

import (
	"fmt"
)

//Problem: https://leetcode.com/problems/minimum-depth-of-binary-tree/

func main() {
	root := &TreeNode{
		Val:  2,
		Left: nil,
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:  4,
				Left: nil,
				Right: &TreeNode{
					Val:  5,
					Left: nil,
					Right: &TreeNode{
						Val:   6,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}
	fmt.Println(minDepth(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minimum(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}
	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}
	return 1 + minimum(minDepth(root.Left), minDepth(root.Right))
}
