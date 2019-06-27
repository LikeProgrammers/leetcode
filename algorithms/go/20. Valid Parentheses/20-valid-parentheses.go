package main

import "fmt"

func main() {
	ss := []string{"()", "()[]{}", "(]", "([)]", "{[]}", ")", "("}
	for _, s := range ss {
		ret := isValid(s)
		fmt.Println(ret)
	}
}

func isValid(s string) bool {
	bracket_map := map[string]string{"(": ")", "[": "]", "{": "}"}
	var brackets []string
	ls := len(s)
	for i := 0; i < ls; i++ {
		v := s[i : i+1]
		if v == "(" || v == "[" || v == "{" {
			brackets = append(brackets, v)
		} else if v == ")" || v == "]" || v == "}" {
			lb := len(brackets)
			if lb <= 0 {
				return false
			}
			end := lb - 1
			if bracket_map[brackets[end]] != v {
				return false
			}
			brackets = brackets[:end]
		} else {
			return false
		}
	}

	if len(brackets) != 0 {
		return false
	}

	return true
}
