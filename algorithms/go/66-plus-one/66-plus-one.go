package main

import (
	"fmt"
)

func main() {
	digits := []int{9, 9}
	ret := plusOne(digits)
	fmt.Println(ret)
}

func plusOne(digits []int) []int {
	if len(digits) < 1 {
		digits = append(digits, 0)
	}

	digits[len(digits)-1] += 1

	ret := []int{}
	carry := 0
	for i := len(digits) - 1; i >= 0; i-- {
		current := digits[i] + carry
		carry = current / 10
		digits[i] = current % 10
		if carry <= 0 {
			break
		}
	}
	if carry > 0 {
		ret = append(ret, carry)
	}

	ret = append(ret, digits...)

	return ret
}
