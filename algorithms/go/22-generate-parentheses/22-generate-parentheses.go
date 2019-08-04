package main

import "fmt"

func main() {
	n := 3
	ret := generateParenthesis(n)
	fmt.Println(ret)
	return
}

var ss = []string{}

func generateParenthesis(n int) []string {
	ss = nil
	s := []byte{}
	l := 0
	r := 0
	backtrack(s, l, r, n)
	return ss
}

func backtrack(s []byte, l int, r int, n int) {
	if len(s) == n*2 {
		ss = append(ss, string(s))
		return
	}
	if l < n {
		sl := make([]byte, len(s))
		copy(sl, s)
		sl = append(sl, '(')
		backtrack(sl, l+1, r, n)
	}
	if r < l {
		sr := make([]byte, len(s))
		copy(sr, s)
		sr = append(sr, ')')
		backtrack(sr, l, r+1, n)
	}
}
