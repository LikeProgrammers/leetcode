package main

import (
	"fmt"
)

func main() {
	word := "Hello"
	ret := detectCapitalUse(word)
	fmt.Println(ret)
}

func detectCapitalUse(word string) bool {
	if len(word) <= 0 {
		return true
	}
	all_capital := true
	all_letter := true
	first := isCapital(rune(word[0]))
	for i, v := range word {
		if !isAlpha(v) {
			all_capital = false
			all_letter = false
			first = false
		}
		if i <= 0 {
			continue
		}
		if isCapital(v) {
			all_letter = false
		}
		if isLetter(v) {
			all_capital = false
		}
	}
	return all_letter || (first && all_capital)
}

func isCapital(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

func isLetter(c rune) bool {
	return c >= 'a' && c <= 'z'
}

func isAlpha(c rune) bool {
	return isCapital(c) || isLetter(c)
}
