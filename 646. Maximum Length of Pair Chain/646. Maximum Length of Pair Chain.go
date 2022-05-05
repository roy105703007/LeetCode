func findLongestChain(pairs [][]int) int {
	dp := make([]int, len(pairs))
	max := 1
	quicksort(pairs)
	for i := 0; i < len(pairs); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if pairs[j][1] < pairs[i][0] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				if dp[i] > max {
					max = dp[i]
				}
			}
		}
	}
	return max
}

func quicksort(a [][]int) [][]int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := 0

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i][0] < a[right][0] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}