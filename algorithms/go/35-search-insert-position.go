package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 5, 6}
	target_arr := []int{5, 2, 7, 0}
	for i := 0; i < len(target_arr); i++ {
		target := target_arr[i]
		ret := searchInsert(nums, target)
		fmt.Println(ret)
	}
}

func searchInsert(nums []int, target int) int {
	ln := len(nums)
	left := 0
	right := ln - 1
	if target < nums[left] {
		return left
	} else if target > nums[right] {
		return right + 1
	}

	for left < right {
		if right-left <= 1 {
			if target == nums[left] {
				return left
			} else if target == nums[right] {
				return right
			} else {
				return right
			}
		}
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid
		} else if target < nums[mid] {
			right = mid
		}
	}

	return 0
}
