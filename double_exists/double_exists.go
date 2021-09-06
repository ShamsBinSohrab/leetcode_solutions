package main

import "fmt"

func main() {
	arr := []int{-2, 0, 10, -19, 4, 6, -8}
	fmt.Println(checkIfExist(arr))
}

func checkIfExist(arr []int) bool {
	m := make(map[int]bool)
	for _, v := range arr {
		if _, ok := m[v*2]; ok {
			return ok
		}
		if v%2 == 0 {
			if _, ok := m[v/2]; ok {
				return ok
			}
		}
		m[v] = true
	}
	return false
}
