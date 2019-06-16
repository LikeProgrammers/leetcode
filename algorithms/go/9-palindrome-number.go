package main

import (
	"fmt"
)

func main() {
	x_arr := []int{121, -121, 0, 10}
	for _, x := range x_arr {
		ret := isPalindrome(x)
		fmt.Println(ret)
	}
}

func isPalindrome(x int) bool {
	ret := false
	sx := x

	if sx < 0 {
		return ret
	}

	// convert source int x to array
	var x_arr []int
	for ; sx > 0; sx /= 10 {
		residue := sx % 10
		x_arr = append(x_arr, residue)
	}

	// convert array to revesal int rx
	lx := len(x_arr)
	rx := 0
	for i := 0; i < lx; i++ {
		num := 1
		for n := 0; n < lx-i-1; n++ {
			num *= 10
		}
		rx += x_arr[i] * num
	}

	// output
	if rx == x {
		ret = true
	} else {
		ret = false
	}

	return ret
}
