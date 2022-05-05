func uniquePaths(m int, n int) int {
	// dp[m][n] = dp[m-1][n] + dp[m][n-1]
	dp := make([]int, m*n)
	dp[0] = 1
	for i := 0; i < m*n; i++ {
		if i < n || i%n == 0 {
			dp[i] = 1
			continue
		}
		dp[i] = dp[i-1] + dp[i-n]
	}
	return dp[m*n-1]
}