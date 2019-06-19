package main

import (
	"fmt"
)

func main() {
	// strs := []string{"flower", "flow", "flight"}
	// strs := []string{"dog", "racecar", "car"}
	strs := []string{"aca", "cba"}
	ret := longestCommonPrefix(strs)
	fmt.Println(ret)
}

func longestCommonPrefix(strs []string) string {
	ret := ""
	var pre_arr []byte
	ls := len(strs)
	if ls <= 0 {
		return ret
	}
	min_len := len(strs[0])
	for i := 0; i < ls; i++ {
		l := len(strs[i])
		if l < min_len {
			min_len = l
		}
	}
	for i := 0; i < min_len; i++ {
		v := strs[0][i]
		j := 0
		for ; j < ls; j++ {
			if strs[j][i] != v {
				break
			}
		}
		if j >= ls {
			pre_arr = append(pre_arr, v)
		} else {
			break
		}
	}

	ret = string(pre_arr)

	return ret
}
