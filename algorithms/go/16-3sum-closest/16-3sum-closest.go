package main

import (
	"fmt"
	"sort"
)

func main() {
	nums_arr := [][]int{{-1, 2, 1, -4}, {0, 0, 0, 0}}
	for _, nums := range nums_arr {
		ret := threeSumClosest(nums, 1)
		fmt.Println(ret)
	}
}

func threeSumClosest(nums []int, target int) int {
	const MAX_INT = 1 << 31
	ln := len(nums)
	ret := 0
	if ln < 3 {
		return ret
	}
	sort.Ints(nums)
	diff := MAX_INT
	for i := 0; i < ln-2; i++ {
		if i-1 >= 0 && nums[i-1] == nums[i] {
			continue
		}
		l := i + 1
		h := ln - 1
		for l < h {
			s := nums[i] + nums[l] + nums[h]
			if s == target {
				ret = target
				return ret
			} else {
				current_diff := abs(s - target)
				if current_diff < diff {
					diff = current_diff
					ret = s
				} else if s < target {
					for l < h && l+1 < ln && nums[l+1] == nums[l] {
						l++
					}
					l++
				} else {
					for l < h && h-1 >= 0 && nums[h-1] == nums[h] {
						h--
					}
					h--
				}
			}
		}
	}

	return ret
}

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -1 * x
	}
}
