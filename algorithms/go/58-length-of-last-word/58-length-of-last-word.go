package main

import (
	"fmt"
)

func main() {
	s := "Hello World "
	ret := lengthOfLastWord(s)
	fmt.Println(ret)
}

func lengthOfLastWord(s string) int {
	ls := len(s)
	if ls <= 0 {
		return 0
	}

	end_index := ls - 1
	for ; end_index >= 0 && s[end_index] == ' '; end_index-- {
		continue
	}

	length := 0
	for i := end_index; i >= 0 && s[i] != ' '; i-- {
		length++
	}

	return length
}
