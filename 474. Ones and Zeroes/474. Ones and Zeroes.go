func findMaxForm(strs []string, m int, n int) int {
	// dp[i] = max(dp[i], dp[i-item] + 1)
	var s0, s1 int
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < len(strs); i++ {
		s0, s1 = count01(strs[i])
		for j := m; j >= s0; j-- {
			for k := n; k >= s1; k-- {
				dp[j][k] = max(dp[j][k], dp[j-s0][k-s1]+1)
			}
		}
	}
	return dp[m][n]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func count01(strs string) (int, int) {
	s0, s1 := 0, 0
	for i := 0; i < len(strs); i++ {
		if (strs[i] - '0') == 0 {
			s0++
		} else {
			s1++
		}
	}
	return s0, s1
}