func maxProfit(prices []int) int {
	//buy[i] = 第i天是buy或colddown的手頭上總金額
	//sell[i] = 第i天是sell或colddown的手頭上總金額
	if len(prices) == 1 {
		return 0
	}
	buy := make([]int, len(prices)+1)
	sell := make([]int, len(prices)+1)
	buy[0] = -prices[0]
	sell[0] = 0
	buy[1] = max(buy[0], -prices[1])
	sell[1] = 0
	for i := 2; i < len(prices)+1; i++ {
		buy[i] = max(sell[i-2]-prices[i-1], buy[i-1])
		sell[i] = max(buy[i-1]+prices[i-1], sell[i-1])
	}
	return sell[len(prices)]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// buy:-1 -1 -1 -1 1
// sell:0 0 1 2 2