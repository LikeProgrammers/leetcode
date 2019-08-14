package main

import (
	"fmt"
)

func main() {
	n := 8
	ret := arrangeCoins(n)
	fmt.Println(ret)
}

func arrangeCoins(n int) int {
	i := 1
	for ; n > i; i++ {
		n -= i
	}
	return i-1
}
