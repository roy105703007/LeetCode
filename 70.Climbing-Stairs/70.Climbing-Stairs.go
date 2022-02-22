func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	pre1, pre2 := 2, 1
	cur := 0
	for i := 3; i <= n; i++ {
		cur = pre1 + pre2
		pre2 = pre1
		pre1 = cur
	}
	return cur
}