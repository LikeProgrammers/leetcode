package main

import (
	"fmt"
)

func main() {
	x := 120

	ret := reverse(x)
	fmt.Println(ret)
}

func reverse(x int) int {
	symbol := true
	if x < 0 {
		x *= -1
		symbol = false
	}
	x_arr := make([]int, 32)
	n := 0
	for n = 0; x > 0; n++ {
		residue := x % 10
		x_arr[n] = residue
		x /= 10
	}

	rx := 0
	for i := 0; i < n; i++ {
		current := 1
		num := n - 1 - i
		for j := 0; j < num; j++ {
			current *= 10
		}
		rx = rx + x_arr[i]*current
	}
	if !symbol {
		rx *= -1
	}

	if rx < (1<<31*-1) || rx > (1<<31-1) {
		return 0
	}

	return rx
}
