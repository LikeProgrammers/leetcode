package main

import (
	"fmt"
)

func main() {
	// digits := "23"
	digits := ""
	ret := letterCombinations(digits)
	fmt.Println(ret)
}

// three solutions
// recursive
var ret []string

func letterCombinations(digits string) []string {
	ret = nil // global variable clear
	ld := len(digits)
	if ld <= 0 {
		return ret
	}

	appendLetter("", digits)

	return ret
}

func appendLetter(combination string, digits string) {
	var letter_map = map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}

	if len(digits) <= 0 {
		ret = append(ret, combination)
		return
	}

	digit := digits[0]
	letter := letter_map[digit]

	for i := 0; i < len(letter); i++ {
		appendLetter(combination+letter[i:i+1], digits[1:])
	}
}

// loop
func letterCombinationsThree(digits string) []string {
	var ret []string
	ld := len(digits)
	if ld <= 0 {
		return ret
	}

	var letter_map = map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}

	for i := 0; i < ld; i++ {
		letters := letter_map[digits[i]]
		var r []string
		for d := 0; d < len(letters); d++ {
			if len(ret) <= 0 {
				r = append(r, letters[d:d+1])
			}
			for j := 0; j < len(ret); j++ {
				r = append(r, ret[j]+letters[d:d+1])
			}
		}
		ret = r
	}

	return ret
}

// queue
