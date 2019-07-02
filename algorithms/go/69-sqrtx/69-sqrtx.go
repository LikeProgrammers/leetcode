package main

import (
	"fmt"
)

func main() {
	x := 2147395599
	ret := mySqrt(x)
	fmt.Println(ret)
}

func mySqrt(x int) int {
	if x <= 0 {
		return 0
	}

	left := 0
	right := x
	for left <= right {
		mid := (left + right) / 2
		pow := int64(mid) * int64(mid)
		if pow == int64(x) {
			return mid
		} else if pow < int64(x) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left - 1
}
