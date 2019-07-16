package intersection

func intersection(nums1 []int, nums2 []int) []int {
	arr := []int{}
	m1 := map[int]bool{}
	m2 := map[int]bool{}
	for _, v := range nums1 {
		m1[v] = true
	}
	for _, v := range nums2 {
		m2[v] = true
	}
	for k, _ := range m1 {
		if _, found := m2[k]; found {
			arr = append(arr, k)
		}
	}

	return arr
}
