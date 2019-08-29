package main

import (
	"fmt"
	"strconv"
)

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	ret := compress(chars)
	fmt.Println(ret, chars)
}

func compress(chars []byte) int {
	cc := []byte{}
	count := 0
	var current byte
	for _, v := range chars {
		if current == v {
			count++
		} else {
			if count > 1 {
				count_str := strconv.Itoa(count)
				cc = append(cc, []byte(count_str)...)
			}
			cc = append(cc, v)
			current = v
			count = 1
		}
	}
	if count > 1 {
		count_str := strconv.Itoa(count)
		cc = append(cc, []byte(count_str)...)
	}
	copy(chars, cc)
	return len(cc)
}

func compress2(chars []byte) int {
	l := len(chars)
	if l < 1 {
		return 0
	}
	cur, total, curChar := 0, 1, chars[0]
	for i := 1; i <= l; i++ {
		if i < l && chars[i] == chars[i-1] {
			total++
			continue
		}
		chars[cur] = curChar
		cur++
		if total > 1 {
			num := strconv.Itoa(total)
			for j := 0; j < len(num); j++ {
				chars[cur] = num[j]
				cur++
			}
		}
		if i < l {
			curChar = chars[i]
		}
		total = 1
	}
	return cur
}
