package main

import (
	"fmt"
)

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	ret := convert(s, numRows)
	fmt.Println(ret) // PAHNAPLSIIGYIR
}

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	ls := len(s)
	z := make([][]byte, numRows)
	// init
	for i := 0; i < numRows; i++ {
		z[i] = make([]byte, ls)
	}

	stepi, stepj := 1, 0
	for k, i, j := 0, -1, 0; k < ls; k++ {
		i += stepi
		j += stepj
		z[i][j] = s[k]
		if i >= numRows-1 {
			stepi = -1
			stepj = 1
		} else if i <= 0 {
			stepi = 1
			stepj = 0
		}
	}

	zs := make([]byte, 0)
	for i := 0; i < len(z); i++ {
		for j := 0; j < len(z[i]); j++ {
			if z[i][j] != 0 {
				zs = append(zs, z[i][j])
			}
		}
	}

	// fmt.Println(len(zs), zs)
	return string(zs)
}
