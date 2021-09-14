package main

import "fmt"

//Problems: https://leetcode.com/problems/pascals-triangle-ii/

func main() {
	fmt.Println(getRow(5))
}

func getRow(rowIndex int) []int {
	triangle := [][]int{{1}, {1, 1}}
	for i := 2; i <= rowIndex; i++ {
		triangle = append(triangle, make([]int, i+1))
		triangle[i][0], triangle[i][i] = 1, 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	return triangle[rowIndex]
}
