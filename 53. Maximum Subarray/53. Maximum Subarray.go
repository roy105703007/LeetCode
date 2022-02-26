func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum := nums[0]
	res := 0
	for i := range nums {
		res += nums[i]
		if res > maxSum {
			maxSum = res
		}
		if res < 0 {
			res = 0
		}
	}
	return maxSum
}