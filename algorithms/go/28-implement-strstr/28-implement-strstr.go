package main

import (
	"fmt"
)

func main() {
	haystack := "aaaaa"
	needle := "bba"
	ret := strStr(haystack, needle)
	fmt.Println(ret)
}

func strStr(haystack string, needle string) int {
	lh := len(haystack)
	ln := len(needle)
	if lh <= 0 && ln <= 0 {
		return 0
	} else if lh <= 0 {
		return -1
	} else if ln <= 0 {
		return 0
	} else if lh < ln {
		return -1
	}

	index := -1
	for i := 0; i < lh-ln+1; i++ {
		str := haystack[i : i+ln]
		if str == needle {
			index = i
			break
		}
	}

	return index
}
