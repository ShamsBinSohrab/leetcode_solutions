package main

//Problem: https://leetcode.com/problems/range-sum-query-immutable/

func main() {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	obj.SumRange(0, 2)
	obj.SumRange(2, 5)
	obj.SumRange(0, 5)
}

type NumArray struct {
	array []int
}

/*func Constructor(nums []int) NumArray {
	return NumArray{array: nums}
}

func (na *NumArray) SumRange(left int, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += na.array[i]
	}
	return sum
}*/

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{array: []int{}}
	}
	temp := make([]int, len(nums))
	temp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		temp[i] = temp[i-1] + nums[i]
	}
	return NumArray{array: temp}
}

func (na *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return na.array[right]
	}
	return na.array[right] - na.array[left-1]
}
