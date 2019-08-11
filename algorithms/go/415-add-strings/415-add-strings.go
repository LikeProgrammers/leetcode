package main

import (
	"fmt"
)

func main() {
	num1 := "456"
	num2 := "789"
	ret := addStrings(num1, num2)
	fmt.Println(ret)
}

func addStrings(num1 string, num2 string) string {
	i := len(num1) - 1
	j := len(num2) - 1
	carry := 0
	num := []byte{}

	for i >= 0 || j >= 0 || carry != 0 {
		sum := carry
		if i >= 0 {
			sum += int(num1[i]) - '0'
			i--
		}
		if j >= 0 {
			sum += int(num2[j]) - '0'
			j--
		}
		num = append(num, byte(sum%10+'0'))
		carry = sum / 10
	}

	for k := 0; k < len(num)/2; k++ {
		num[k], num[len(num)-1-k] = num[len(num)-1-k], num[k]
	}

	return string(num)
}
