package main

import "fmt"

//Problems: https://leetcode.com/problems/count-primes/

func main() {
	fmt.Println(countPrimes(10))
}

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	var count int
	notPrimes := make([]bool, n)
	notPrimes[0], notPrimes[1] = true, true
	for i := 2; i < n; i++ {
		if notPrimes[i] {
			continue
		}
		count++
		for j := i * 2; j < n; j += i {
			notPrimes[j] = true
		}
	}
	return count
}
