package main

import (
	"fmt"
)

//Problem: https://leetcode.com/problems/same-tree/

func main() {
	p := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	q := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(isSameTree(p, q))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(node *TreeNode, values *[]int) {
	if node == nil {
		*values = append(*values, -1)
		return
	}
	*values = append(*values, node.Val)
	traverse(node.Left, values)
	traverse(node.Right, values)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var pValues, qValues = make([]int, 0), make([]int, 0)
	traverse(p, &pValues)
	traverse(q, &qValues)
	if len(pValues) == len(qValues) {
		for i := 0; i < len(pValues); i++ {
			if pValues[i] != qValues[i] {
				return false
			}
		}
		return true
	}
	return false
}
