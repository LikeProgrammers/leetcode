package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums := []int{}
	i, j := 0, 0
	for i < m || j < n {
		if i < m && j < n {
			if nums1[i] <= nums2[j] {
				nums = append(nums, nums1[i])
				i++
			} else {
				nums = append(nums, nums2[j])
				j++
			}
		} else if i < m {
			nums = append(nums, nums1[i])
			i++
		} else if j < n {
			nums = append(nums, nums2[j])
			j++
		} else {
			break
		}
	}

	i, j = 0, 0
	for j < len(nums) {
		nums1[i] = nums[j]
		i++
		j++
	}
}
