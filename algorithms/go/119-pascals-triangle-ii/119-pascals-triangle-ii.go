package triangle_ii

func getRow(rowIndex int) []int {
	list := []int{}
	pre_list := []int{}

	if rowIndex >= 0 {
		list = append(list, 1)
	}

	if rowIndex >= 1 {
		list = append(list, 1)
		pre_list = list
	}

	for i := 2; i <= rowIndex; i++ {
		list = nil
		list = append(list, 1)
		for j := 0; j < len(pre_list)-1; j++ {
			list = append(list, pre_list[j]+pre_list[j+1])
		}
		list = append(list, 1)
		pre_list = list
	}

	return list
}
