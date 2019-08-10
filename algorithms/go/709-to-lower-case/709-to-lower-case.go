package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello"
	ret := toLowerCase(str)
	fmt.Println(ret)
}

func toLowerCase(str string) string {
	return strings.ToLower(str)
}
