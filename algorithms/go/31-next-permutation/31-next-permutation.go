package nextPermutation

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		swap(nums, i, j)
	}
	reverse(nums, i+1)
}

func reverse(nums []int, start int) {
	i := start
	j := len(nums) - 1
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func swap(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
