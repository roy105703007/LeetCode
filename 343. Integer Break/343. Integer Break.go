func integerBreak(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		for j := 1; j < i/2+1; j++ {
			if i != n {
				dp[i] = max(max(i, dp[j]*dp[i-j]), dp[i])
			} else {
				dp[i] = max(dp[i], dp[j]*dp[i-j])
			}
		}
	}
	return dp[n]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}