package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 15
	ret := fizzBuzz(n)
	fmt.Println(ret)
}

func fizzBuzz(n int) []string {
	ret := []string{}
	for i := 1; i <= n; i++ {
		remainder3 := i % 3
		remainder5 := i % 5
		if remainder3 == 0 && remainder5 == 0 {
			ret = append(ret, "FizzBuzz")
		} else if remainder3 == 0 {
			ret = append(ret, "Fizz")
		} else if remainder5 == 0 {
			ret = append(ret, "Buzz")
		} else {
			ret = append(ret, strconv.Itoa(i))
		}
	}
	return ret
}
