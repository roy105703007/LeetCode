func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	minItem := min(nums[left], nums[right])
	mid := 0
	for left < right-1 {
		mid = left + (right-left)/2
		if nums[mid] > nums[left] {
			left = mid
		} else {
			minItem = min(minItem, nums[mid])
			right = mid
		}
	}
	return minItem
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}