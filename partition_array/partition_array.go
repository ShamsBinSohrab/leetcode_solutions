package main

import "fmt"

func main() {
	arr := []int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}
	fmt.Println(canThreePartsEqualSum(arr))
}

func canThreePartsEqualSum(arr []int) bool {
	var sum int
	for _, n := range arr {
		sum += n
	}
	if sum%3 == 0 {
		sum /= 3
		var part, cs int
		for _, n := range arr {
			cs += n
			if cs == sum {
				part += 1
				cs = 0
			}
		}
		return part >= 3
	}
	return false
}
