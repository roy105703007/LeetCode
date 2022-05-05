func numDecodings(s string) int {
	if len(s) == 1 {
		if (s[0] - '0') == 0 {
			return 0
		}
		return 1
	}
	if len(s) == 2 {
		if (s[0]-'0') == 0 || ((s[0]-'0') > 2 && (s[1]-'0') == 0) {
			return 0
		} else if (s[0]-'0')*10+(s[1]-'0') > 26 || (s[1]-'0') == 0 {
			return 1
		} else {
			return 2
		}
	}
	dp := make([]int, len(s))
	dp[0] = 1
	if (s[0]-'0') == 0 || ((s[0]-'0') > 2 && (s[1]-'0') == 0) {
		return 0
	} else if (s[0]-'0')*10+(s[1]-'0') > 26 || (s[1]-'0') == 0 {
		dp[1] = 1
	} else {
		dp[1] = 2
	}
	for i := 2; i < len(s); i++ { //dp[i] = dp[i-1]沒綁一起 + dp[i-2]綁一起
		if s[i]-'0' == 0 {
			if s[i-1]-'0' < 3 && s[i-1]-'0' > 0 {
				dp[i] = dp[i-2]
			} else {
				return 0
			}
		} else {
			dp[i] += dp[i-1]
			if s[i-1]-'0' != 0 && (s[i-1]-'0')*10+(s[i]-'0') <= 26 {
				dp[i] += dp[i-2]
			}
		}

	}
	return dp[len(s)-1]
}