func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	left := 0
	right := len(nums) - 1
	mid := (left + right) / 2
	for left < right {
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[left] {
			if target < nums[left] {
				left = mid
			} else if target < nums[mid] {
				right = mid
			} else {
				left = mid
			}
		} else {
			if target < nums[mid] {
				right = mid
			} else if target < nums[left] {
				left = mid
			} else {
				right = mid
			}
		}
		mid = (left + right) / 2
	}
	return -1
}