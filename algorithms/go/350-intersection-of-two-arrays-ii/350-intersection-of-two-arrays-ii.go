package intersect

func intersect(nums1 []int, nums2 []int) []int {
	m := []int{}

	m1 := map[int]int{}
	m2 := map[int]int{}

	for _, v := range nums1 {
		m1[v]++
	}
	for _, v := range nums2 {
		m2[v]++
	}

	for v, c1 := range m1 {
		if c2, found := m2[v]; found {
			c := minInt(c1, c2)
			for i := 0; i < c; i++ {
				m = append(m, v)
			}
		}
	}

	return m
}

func minInt(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
