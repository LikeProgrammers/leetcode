package main

import (
	"fmt"
)

func main() {
	a := "001010"
	b := "001011"
	ret := addBinary(a, b)
	fmt.Println(ret)
}

func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)

	lmax := maxInt(la, lb)
	sum_arr := make([]byte, lmax)
	ia := la - 1
	ib := lb - 1
	is := lmax - 1
	carry := 0
	for is >= 0 {
		current := 0
		if ia >= 0 {
			current += int(a[ia]) - '0'
		}
		if ib >= 0 {
			current += int(b[ib]) - '0'
		}
		current += carry
		carry = current / 2
		sum_arr[is] = byte(current%2 + '0')
		is--
		ia--
		ib--
	}

	sum := []byte{}
	if carry > 0 {
		sum = append(sum, byte(carry)+'0')
	}
	sum = append(sum, sum_arr...)

	i := 0
	for ; i < len(sum); i++ {
		if sum[i] > '0' {
			break
		}
	}

	if i >= len(sum) {
		return "0"
	} else {
		return string(sum[i:])
	}
}

func maxInt(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
