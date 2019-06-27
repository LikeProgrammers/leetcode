package main

import "fmt"

func main() {

	nums := []int{2, 7, 11, 15}
	target := 9
	var ret = twoSum(nums, target)

	fmt.Println(ret) // [0, 1]
}

func twoSum(nums []int, target int) []int {

	nums_len := len(nums)

	if nums_len < 2 {
		return nil
	}

	for i := 0; i < nums_len-1; i++ {
		diff := target - nums[i]
		for j := i + 1; j < nums_len; j++ {
			if nums[j] == diff {
				return []int{i, j}
			}
		}
	}

	return nil
}
