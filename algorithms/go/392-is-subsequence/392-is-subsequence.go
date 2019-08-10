package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := "abc"
	t := "ahbgd"
	ret := isSubsequence(s, t)
	fmt.Println(ret)
}

func isSubsequence1(s string, t string) bool {
	// [a-z]*a[a-z]*b[a-z]*c[a-z]*
	ret := false
	gap := "[a-z]*"
	pattern := []string{}
	for _, v := range s {
		pattern = append(pattern, gap)
		pattern = append(pattern, string(v))
	}
	if len(pattern) != 0 {
		pattern = append(pattern, gap)
	}
	patterns := strings.Join(pattern, "")

	matched, err := regexp.MatchString(patterns, t)
	if matched && err == nil {
		ret = true
	}
	return ret
}

func isSubsequence(s string, t string) bool {
	ls := len(s)
	if ls == 0 {
		return true
	}
	i, j := 0, 0
	for lt := len(t); j < lt; j++ {
		if s[i] == t[j] {
			if i++; i == ls {
				return true
			}
		}
	}
	return false
}
