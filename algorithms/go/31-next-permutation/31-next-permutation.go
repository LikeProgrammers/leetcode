package nextPermutation

import "sort"

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			nums[i], nums[i-1] = nums[i-1], nums[i]
			return
		}
	}
	sort.Ints(nums)
	return
}
