package main

import "fmt"

func main() {

	s_arr := []string{"abcabcbb", "bbbbb", "pwwkew", "abcad", "tmmzuxt"}
	// s_arr := []string{"tmmzuxt"}
	for _, s := range s_arr {
		var ret = lengthOfLongestSubstring(s)
		fmt.Println(ret) // 3, 1, 3, 4, 5
	}
}

func lengthOfLongestSubstring(s string) int {
	max_len, sub_len, start_i := 0, 0, 0

	for k, v := range s {
		in_sub := false
		sub_i := 0
		sub_s := s[start_i:k]
		for kk, vv := range sub_s {
			if vv == v {
				in_sub = true
				sub_i = kk
				break
			}
		}
		if in_sub {
			if sub_len > max_len {
				max_len = sub_len
			}
			start_i += sub_i + 1
			sub_len = k - start_i + 1
		} else {
			sub_len++
		}
	}

	if sub_len > max_len {
		max_len = sub_len
	}

	return max_len
}
