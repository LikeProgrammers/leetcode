package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums := []int{1, 0, -1, 0, -2, 2}
	// target := 0
	nums := []int{2, 1, 0, -1}
	target := 2
	ret := fourSum(nums, target)
	fmt.Println(ret)
}

func fourSum1(nums []int, target int) [][]int {
	sort.Ints(nums)
	ln := len(nums)
	msum := [][]int{}
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if nums[i]+nums[ln-1]+nums[ln-2]+nums[ln-3] < target {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			if nums[i]+nums[j]+nums[ln-1]+nums[ln-2] < target {
				continue
			}

			for k := j + 1; k < len(nums)-1; k++ {
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}
				if nums[i]+nums[j]+nums[k]+nums[k+1] > target {
					break
				}
				if nums[i]+nums[j]+nums[k]+nums[ln-1] < target {
					continue
				}

				for l := k + 1; l < len(nums); l++ {
					if l > k+1 && nums[l] == nums[l-1] {
						continue
					}
					if nums[i]+nums[j]+nums[k]+nums[l] > target {
						break
					}
					if nums[i]+nums[j]+nums[k]+nums[l] < target {
						continue
					}
					if nums[i]+nums[j]+nums[k]+nums[l] == target {
						lsum := []int{nums[i], nums[j], nums[k], nums[l]}
						msum = append(msum, lsum)
					}
				}
			}
		}
	}
	return msum
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ln := len(nums)
	msum := [][]int{}
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if nums[i]+nums[ln-1]+nums[ln-2]+nums[ln-3] < target {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			if nums[i]+nums[j]+nums[ln-1]+nums[ln-2] < target {
				continue
			}

			l := j + 1
			h := ln - 1
			for l < h {
				s := nums[i] + nums[j] + nums[l] + nums[h]
				if s == target {
					r := []int{nums[i], nums[j], nums[l], nums[h]}
					msum = append(msum, r)

					for l < h && l+1 < ln && nums[l+1] == nums[l] {
						l++
					}
					for l < h && h-1 >= 0 && nums[h-1] == nums[h] {
						h--
					}
					l++
					h--
				} else if s < target {
					l++
				} else {
					h--
				}
			}
		}
	}
	return msum
}
