package main

import "fmt"

//Problem: https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	ans := sortedArrayToBST(nums)
	fmt.Println(ans)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if mid := len(nums) / 2; len(nums) != 0 {
		return &TreeNode{
			Val:   nums[mid],
			Left:  sortedArrayToBST(nums[:mid]),
			Right: sortedArrayToBST(nums[mid+1:]),
		}
	}
	return nil
}
