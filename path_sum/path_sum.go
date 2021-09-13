package main

import (
	"fmt"
)

//Problem: https://leetcode.com/problems/path-sum/

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  4,
				Left: nil,
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	fmt.Println(root)
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: nil,
			},
			Right: nil,
		},
		Right: nil,
	}
	fmt.Println(root2)
	fmt.Println(hasPathSum(root2, 6))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	return (root.Left == nil && root.Right == nil && targetSum == 0) || hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}
