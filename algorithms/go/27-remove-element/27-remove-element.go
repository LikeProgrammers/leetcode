package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	ret := removeElement(nums, val)
	fmt.Println(ret, nums)
}

func removeElement(nums []int, val int) int {
	n := 0
	ln := len(nums)
	for i := 0; i < ln; i++ {
		if nums[i] != val {
			if n < i {
				nums[n] = nums[i]
			}
			n++
		}
	}

	return n
}
