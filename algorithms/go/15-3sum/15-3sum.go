package main

import (
	"fmt"
	"sort"
)

func main() {
	nums_arr := [][]int{{-1, 0, 1, 2, -1, -4}, {0, 0, 0, 0}}
	for _, nums := range nums_arr {
		ret := threeSum(nums)
		fmt.Println(ret)
	}
}

func threeSum(nums []int) [][]int {
	ln := len(nums)
	var ret [][]int = nil
	if ln < 3 {
		return ret
	}
	sort.Ints(nums)
	for i := 0; i < ln-2; i++ {
		if i-1 >= 0 && nums[i-1] == nums[i] {
			continue
		}
		l := i + 1
		h := ln - 1
		for l < h {
			s := nums[i] + nums[l] + nums[h]
			if s == 0 {
				r := []int{nums[i], nums[l], nums[h]}
				ret = append(ret, r)

				for l < h && l+1 < ln && nums[l+1] == nums[l] {
					l++
				}
				for l < h && h-1 >= 0 && nums[h-1] == nums[h] {
					h--
				}
				l++
				h--
			} else if s < 0 {
				l++
			} else {
				h--
			}
		}
	}

	return ret
}

// repeated combination
func threeSumRe(nums []int) [][]int {
	ln := len(nums)
	var ret [][]int = nil
	if ln < 3 {
		return ret
	}
	for i := 0; i < ln-2; i++ {
		for j := i + 1; j < ln-1; j++ {
			for k := j + 1; k < ln; k++ {
				sum := nums[i] + nums[j] + nums[k]
				if sum == 0 {
					r := []int{nums[i], nums[j], nums[k]}
					ret = append(ret, r)
				}
			}
		}
	}

	return ret
}
