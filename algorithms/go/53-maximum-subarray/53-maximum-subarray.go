package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	ret := maxSubArray(nums)
	fmt.Println(ret)
}

func maxSubArray(nums []int) int {
	ln := len(nums)
	if ln <= 0 {
		return 0
	}
	sum := nums[0]
	max := nums[0]
	for i := 1; i < ln; i++ {
		sum = maxInt(sum+nums[i], nums[i])
		max = maxInt(max, sum)
	}

	return max
}

func maxInt(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
