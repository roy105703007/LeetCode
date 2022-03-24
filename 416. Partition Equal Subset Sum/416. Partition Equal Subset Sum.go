func canPartition(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2
	dp := make([][]bool, sum+1)
	for i := range dp {
		dp[i] = make([]bool, len(nums)+1)
	}
	dp[0][0] = true
	for i := 1; i <= len(nums); i++ {
		dp[0][i] = true
	}
	for i := 1; i <= sum; i++ {
		dp[i][0] = false
	}

	for i := 1; i <= sum; i++ {
		for j := 1; j <= len(nums); j++ {
			if nums[j-1] <= i {
				dp[i][j] = dp[i][j-1] || dp[i-nums[j-1]][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[sum][len(nums)]
}