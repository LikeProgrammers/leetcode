package triangle

func generate(numRows int) [][]int {
	tower := [][]int{}
	pre_list := []int{}

	if numRows >= 1 {
		list := []int{1}
		tower = append(tower, list)
	}

	if numRows >= 2 {
		list := []int{1, 1}
		tower = append(tower, list)
		pre_list = list
	}

	for i := 2; i < numRows; i++ {
		list := []int{1}
		for j := 0; j < len(pre_list)-1; j++ {
			list = append(list, pre_list[j]+pre_list[j+1])
		}
		list = append(list, 1)
		tower = append(tower, list)
		pre_list = list
	}

	return tower
}
