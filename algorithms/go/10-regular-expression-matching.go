package main

import (
	"fmt"
)

func main() {
	sp_map := [][]string{{"aa", "a"}, {"aa", "a*"}, {"ab", ".*"}, {"aab", "c*a*b"}, {"mississippi", "mis*is*p*."}}
	for _, sp_arr := range sp_map {
		ret := isMatch(sp_arr[0], sp_arr[1])
		fmt.Println(ret)
	}
}

func isMatch(s string, p string) bool {
	// too hard, talk about it later
}
