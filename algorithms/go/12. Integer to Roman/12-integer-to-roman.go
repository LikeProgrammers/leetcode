package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 58, 54, 1994}
	// nums := []int{9}
	for _, num := range nums {
		ret := intToRoman(num)
		fmt.Println(ret)
	}
}

func intToRoman(num int) string {
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	ret := ""

	if num < 1 || num > 3999 {
		return ret
	}

	roman := ""
	for i := 0; num > 0; i++ {
		for num >= values[i] {
			num -= values[i]
			roman += symbols[i]
		}
	}

	ret = roman
	return ret
}
