func combinationSum4(nums []int, target int) int {
	// dp[i] = sum(dp[i-item])
	dp := make([]int, target+1)
	dp[0] = 1
	for i := range dp {
		for j := range nums {
			if i-nums[j] >= 0 {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}

