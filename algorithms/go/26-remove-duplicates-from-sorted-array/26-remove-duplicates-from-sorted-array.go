package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 2, 3}
	ret := removeDuplicates(nums)
	fmt.Println(ret)
}

func removeDuplicates(nums []int) int {
	n := 0
	ln := len(nums)
	for i := 1; i < ln; i++ {
		if nums[i] != nums[n] {
			n++
			nums[n] = nums[i]
		}
	}

	return n + 1
}
