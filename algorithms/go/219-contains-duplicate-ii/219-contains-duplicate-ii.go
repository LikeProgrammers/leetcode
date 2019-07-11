package containsNearbyDuplicate

func containsNearbyDuplicate(nums []int, k int) bool {
	set := map[int]int{}

	for i, v := range nums {
		if index, found := set[v]; found {
			if i-index <= k {
				return true
			} else {
				set[v] = i
			}
		} else {
			set[v] = i
		}
	}

	return false
}
