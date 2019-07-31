package searchRange

func searchRange(nums []int, target int) []int {
    left := 0
    right := len(nums) - 1
    l, r := -1, -1

    for left <= right {
        i := (left + right) / 2
        if nums[i] == target {
            for l = i; l > 0; l-- {
                if nums[i] != nums[l-1] {
                    break
                }
            }
            for r = i; r < len(nums) - 1; r++ {
                if nums[i] != nums[r+1] {
                    break
                }
            }
            break
        } else if nums[i] > target {
            if right == i {
                right = i - 1
            } else {
                right = i
            }
        } else if nums[i] < target {
            if left == i {
                left = i + 1
            } else {
                left = i
            }
        }
    }

    return []int{l, r}
}