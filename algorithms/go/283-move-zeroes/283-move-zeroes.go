package moveZeroes

// [0,1,0,3,12] [0,0,0,0] [1,2,0,0] [1,2,0,1] [] [1,2,3,4]
func moveZeroes1(nums []int) {
	ln := len(nums)

	i_zero := -1

	for i := 0; i < ln; i++ {
		if i_zero == -1 && nums[i] == 0 {
			i_zero = i
		}
		if i_zero >= 0 && i > i_zero && nums[i] != 0 {
			nums[i_zero] = nums[i]
			i_zero++
		}
	}

	if i_zero >= 0 {
		for i := i_zero; i < ln; i++ {
			nums[i] = 0
		}
	}
}

// Complete first, then optimize
func moveZeroes(nums []int) {
	i_zero := 0
	for _, num := range nums {
		if num != 0 {
			nums[i_zero] = num
			i_zero++
		}
	}
	for ; i_zero < len(nums); i_zero++ {
		nums[i_zero] = 0
	}
}
