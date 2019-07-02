package main

import (
	"fmt"
)

func main() {
	n := 100
	ret := climbStairs(n)
	fmt.Println(ret)
}

func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}

	o := 1
	t := 2
	for i := 2; i < n; i++ {
		s := o + t
		o = t
		t = s
	}

	return t
}
