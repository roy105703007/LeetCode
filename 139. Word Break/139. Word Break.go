func wordBreak(s string, wordDict []string) bool {
	// dp[i+word] = dp[i] + some word, if dp[i] is true
	dp := make([]bool, len(s)+1)
	var flag bool
	var sPoint int
	for i := range dp {
		if i == len(dp)-1 {
			break
		}
		if dp[i] == true || i == 0 {
			for j := range wordDict {
				flag = true
				sPoint = i
				for k := 0; k < len(wordDict[j]); k++ {
					if wordDict[j][k] != s[sPoint] {
						flag = false
						break
					}
					sPoint++
					if sPoint >= len(dp)-1 {
						break
					}
				}
				if flag == true && i+len(wordDict[j]) < len(dp) {
					dp[i+len(wordDict[j])] = true
				}
			}
		}
	}
	return dp[len(dp)-1]
}