package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := "  "
	ret := countSegments(s)
	fmt.Println(ret)
}

func countSegments(s string) int {
	st := strings.TrimSpace(s)
	if len(st) == 0 {
		return 0
	}
	reg := regexp.MustCompile(`\s+`)
	segments := reg.Split(st, -1)
	return len(segments)
}
