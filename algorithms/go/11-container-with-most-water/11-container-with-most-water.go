package main

import (
	"fmt"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	ret := maxArea(height)
	fmt.Println(ret)
}

// O(n)
func maxArea(height []int) int {
	ret := 0
	lh := len(height)
	if lh < 2 {
		return ret
	}

	left := 0
	right := lh - 1
	max_area := 0
	for left < right {
		x, y := 0, 0
		x = right - left
		if height[left] <= height[right] {
			y = height[left]
		} else {
			y = height[right]
		}
		area := x * y
		if area > max_area {
			max_area = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	ret = max_area
	return ret
}

// O(n^2)
func maxAreaBad(height []int) int {
	ret := 0
	lh := len(height)
	if lh < 2 {
		return ret
	}

	max_area := 0
	for i := 0; i < lh-1; i++ {
		for j := i + 1; j < lh; j++ {
			x, y := 0, 0
			if height[i] <= height[j] {
				y = height[i]
			} else {
				y = height[j]
			}
			x = j - i
			area := x * y
			if area > max_area {
				max_area = area
			}
		}
	}

	ret = max_area
	return ret
}
