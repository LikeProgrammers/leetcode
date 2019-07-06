package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 1, 2, 1, 2}
	ret := singleNumber(nums)
	fmt.Println(ret)
}

func singleNumber(nums []int) int {
	ret := 0

	for _, v := range nums {
		ret ^= v
	}

	return ret
}
