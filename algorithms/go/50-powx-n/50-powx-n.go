package main

import (
	"fmt"
)

func main() {
	fmt.Println(myPow(2.00000, 10)) // 1024.00000
	fmt.Println(myPow(2.10000, 3)) // 9.26100
	fmt.Println(myPow(2.00000, -2)) // 0.25000
	fmt.Println(myPow(1.00000, -2147483648)) // 1.0
}

func myPow1(x float64, n int) float64 {
    var p float64 = 1
    var nn int64 = int64(n)

	var i int64 = 0
	var pre float64 = p
    if nn > 0 {
        for i = 0; i < nn; i++ {
			p *= x
			if p == pre {
				break
			}
			pre = p
        }
    } else if nn < 0 {
        nn *= -1
        for i = 0; i < nn; i++ {
			p *= x
			if p == pre {
				break
			}
			pre = p
        }
        p = 1 / p
    }

    return p
}

func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    if n < 0 {
        return 1 / myPow(x, -n)
    }
    y := myPow(x, n / 2)
    if n % 2 == 0 {
        return  y * y
    } else {
        return x * y * y
    }
}
