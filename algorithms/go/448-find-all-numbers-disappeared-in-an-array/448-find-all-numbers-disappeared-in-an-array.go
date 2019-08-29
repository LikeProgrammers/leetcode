package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	ret := findDisappearedNumbers(nums)
	fmt.Println(ret)
}

func findDisappearedNumbers(nums []int) []int {
	for _, v := range nums {
		index := abs(v) - 1
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}
	list := []int{}
	for i, v := range nums {
		if v > 0 {
			list = append(list, i+1)
		}
	}
	return list
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -1 * x
}
