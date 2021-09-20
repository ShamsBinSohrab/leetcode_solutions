package main

import "fmt"

//Problem: https://leetcode.com/problems/binary-tree-postorder-traversal/

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

	fmt.Println(postorderTraversal(root))
}

var postorder func(node *TreeNode)

func postorderTraversal(root *TreeNode) []int {
	var nodes []int
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		nodes = append(nodes, node.Val)
	}
	postorder(root)
	return nodes
}
