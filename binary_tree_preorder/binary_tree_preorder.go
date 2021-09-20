package main

import "fmt"

//Problem: https://leetcode.com/problems/binary-tree-preorder-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

	fmt.Println(preorderTraversal(root))
}

var preorder func(node *TreeNode)

func preorderTraversal(root *TreeNode) []int {
	var nodes []int
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		nodes = append(nodes, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return nodes
}
