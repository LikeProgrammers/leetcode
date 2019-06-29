package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 10
	ret := countAndSay(n)
	fmt.Println(ret)
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	s := countAndSay(n - 1)

	ls := len(s)
	cs := ""
	for i := 0; i < ls; {
		count := 1
		current := s[i]
		for j := i + 1; j < ls && s[j] == current; j++ {
			count++
		}
		cs += strconv.Itoa(count) + string(current)
		i += count
	}

	return cs
}
