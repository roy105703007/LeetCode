func change(amount int, coins []int) int {
	// dp[i] += dp[i-coin] from 0 ... n
	//     1 2 5
	// 0   0 0
	// 1   1 1
	// 2   1 2
	// 3   1 2
	// 4   1 3
	// 5   1 3
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j < len(dp); j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}