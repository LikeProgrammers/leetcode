package main

import (
	"fmt"
)

func main() {
	str_arr := []string{"42", "+42", "-42", "420b2", "   -42", "4193 with words", "words and 987", "-91283472332", "0", "-0", "-04", "04", "9223372036854775808", "  0000000000012345678"}
	// str_arr := []string{"9223372036854775808"}
	for _, str := range str_arr {
		ret := myAtoi(str)
		fmt.Println(ret)
	}
}

func myAtoi(str string) int {
	const MAX_INT int = (1 << 31) - 1
	const MIN_INT int = (1 << 31) * -1
	const MAX_LEN int = 10

	ret := 0
	var x int = 0
	symbol := true
	ls := len(str)

	// find start index and symbol
	start_i := 0
	for start_i = 0; start_i < ls; start_i++ {
		if str[start_i] != ' ' {
			break
		}
	}
	if start_i >= ls {
		return ret
	}
	// find symbol
	if str[start_i] == '-' {
		symbol = false
		start_i++
	} else if str[start_i] == '+' {
		symbol = true
		start_i++
	} else if str[start_i] >= '0' && str[start_i] <= '9' {
		symbol = true
	} else {
		return ret
	}
	if start_i >= ls {
		return ret
	}
	for ; start_i < ls; start_i++ {
		if str[start_i] != '0' {
			break
		}
	}
	if start_i >= ls {
		return ret
	}
	// find end index
	end_i := 0
	for end_i = start_i; end_i < ls; end_i++ {
		if str[end_i] < '0' || str[end_i] > '9' {
			break
		}
	}
	if end_i <= start_i {
		return ret
	}
	// get valid string
	s_arr := str[start_i:end_i]
	lsa := len(s_arr)
	if lsa > MAX_LEN {
		s_arr = s_arr[0 : MAX_LEN+1]
		lsa = len(s_arr)
	}
	// convert to int
	for i := 0; i < lsa; i++ {
		num := 1
		n := lsa - i - 1
		for j := 0; j < n; j++ {
			num *= 10
		}
		x = x + num*(int(s_arr[i])-int('0'))
	}
	if !symbol {
		x *= -1
	}
	// range check
	if x > MAX_INT {
		x = MAX_INT
	} else if x < MIN_INT {
		x = MIN_INT
	}
	// output
	ret = x
	return ret
}
