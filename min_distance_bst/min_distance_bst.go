package main

import (
	"fmt"
	"math"
	"sort"
)

//Problem: https://leetcode.com/problems/minimum-distance-between-bst-nodes/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	_6 := TreeNode{
		Val:   52,
		Left:  nil,
		Right: nil,
	}
	_1 := TreeNode{
		Val:   49,
		Left:  nil,
		Right: &_6,
	}
	_3 := TreeNode{
		Val:   89,
		Left:  nil,
		Right: nil,
	}
	_2 := TreeNode{
		Val:   69,
		Left:  &_1,
		Right: &_3,
	}
	root := TreeNode{
		Val:   90,
		Left:  &_2,
		Right: nil,
	}
	fmt.Println(minDiffInBST(&root))
}
var bfs func(node *TreeNode)

func minDiffInBST(root *TreeNode) int {
	dist := make([]int, 0)
	bfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		bfs(node.Left)
		bfs(node.Right)
		dist = append(dist, node.Val)
	}
	bfs(root)
	sort.Ints(dist)
	min := math.MaxInt32
	for i, v := range dist {
		if (i != len(dist) -1) && (dist[i + 1] - v < min) {
			min = dist[i + 1] - v
		}
	}
	return min
}