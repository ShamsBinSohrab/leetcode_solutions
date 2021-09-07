package main

import "fmt"

//Problem: https://leetcode.com/problems/binary-tree-inorder-traversal/

func main() {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(inorderTraversal(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var traverse func(node *TreeNode)

func inorderTraversal(root *TreeNode) []int {
	values := make([]int, 0)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		values = append(values, node.Val)
		traverse(node.Right)
	}
	traverse(root)
	return values
}
