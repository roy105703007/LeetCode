func moveZeroes(nums []int) {
	count := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[count] = nums[i]
			count += 1
		}
	}
	for i := count; i < len(nums); i++ {
		nums[i] = 0
	}
	return
}