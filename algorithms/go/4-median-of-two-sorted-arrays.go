package main

import "fmt"

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	ret := findMedianSortedArrays(nums1, nums2)
	fmt.Println(ret)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lnums1, lnums2 := len(nums1), len(nums2)
	lnums := lnums1 + lnums2
	nums := make([]int, lnums)
	for i, i1, i2 := 0, 0, 0; i1 < lnums1 || i2 < lnums2; i++ {

		if i1 < lnums1 && i2 < lnums2 {
			if nums1[i1] <= nums2[i2] {
				nums[i] = nums1[i1]
				i1++
			} else {
				nums[i] = nums2[i2]
				i2++
			}
		} else if i1 >= lnums1 {
			nums[i] = nums2[i2]
			i2++
		} else if i2 >= lnums2 {
			nums[i] = nums1[i1]
			i1++
		}
	}

	ret := 0.0
	m := lnums / 2
	if lnums%2 != 0 {
		ret = float64(nums[m])
	} else {
		ret = float64(nums[m]+nums[m-1]) / 2
	}

	return ret
}
