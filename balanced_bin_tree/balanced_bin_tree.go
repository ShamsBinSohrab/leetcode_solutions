package main

import "fmt"

//Problem: https://leetcode.com/problems/balanced-binary-tree/

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
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
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(isBalanced(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var depth func(*TreeNode) int

func isBalanced(root *TreeNode) bool {
	var max int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := depth(node.Left)
		right := depth(node.Right)
		if diff := left - right; diff < 0 {
			if v := diff * -1; max < v {
				max = v
			}
			return 1 + right
		} else if max < diff {
			max = diff
		}
		return 1 + left
	}
	depth(root)
	return max <= 1
}
