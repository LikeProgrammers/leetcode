package rotate

func rotate(nums []int, k int) {
	ln := len(nums)
	r_nums := make([]int, ln)
	for i, v := range nums {
		index := (i + k) % ln
		r_nums[index] = v
	}
	copy(nums, r_nums)
}
