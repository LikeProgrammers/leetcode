package main

import "fmt"

func main() {
	s_arr := []string{"babad", "cbbd", "aaaa", "bacab"}
	for _, s := range s_arr {
		ret := longestPalindrome(s)
		fmt.Println(ret)
	}
}

func longestPalindrome(s string) string {
	var max_s string
	ls := len(s)

	for k, _ := range s {
		left := k - 1
		right := k + 1
		step := 0
		for {
			if left < 0 || right >= ls {
				break
			}
			if s[left] == s[right] {
				step += 1
				left--
				right++
			} else {
				break
			}
		}
		sub_s := s[k-step : k+step+1]
		if len(sub_s) > len(max_s) {
			max_s = sub_s
		}
	}

	for k, _ := range s {
		left := k - 1
		right := k + 2
		step := 0

		if k+1 >= ls {
			break
		}
		if s[k] != s[k+1] {
			continue
		}

		for {
			if left < 0 || right >= ls {
				break
			}
			if s[left] == s[right] {
				step += 1
				left--
				right++
			} else {
				break
			}
		}
		sub_s := s[k-step : k+step+2]
		if len(sub_s) > len(max_s) {
			max_s = sub_s
		}
	}

	return max_s
}
