func wiggleMaxLength(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	if len(nums) == 2 {
		if nums[0] == nums[1] {
			return 1
		}
		return len(nums)
	}
	var pre bool
	res := 2
	if nums[0] < nums[1] {
		pre = true //山峰
		for i := 2; i < len(nums); i++ {
			if pre {
				if nums[i] < nums[i-1] {
					res += 1
					pre = false
				}
			} else {
				if nums[i] > nums[i-1] {
					res += 1
					pre = true
				}
			}
		}
	} else if nums[0] > nums[1] {
		pre = false //山峰
		for i := 2; i < len(nums); i++ {
			if pre {
				if nums[i] < nums[i-1] {
					res += 1
					pre = false
				}
			} else {
				if nums[i] > nums[i-1] {
					res += 1
					pre = true
				}
			}
		}
	} else {
		res = wiggleMaxLength(nums[1:])
	}
	return res
}