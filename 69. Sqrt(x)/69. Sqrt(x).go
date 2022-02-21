func mySqrt(x int) int {
	l := 0
	h := x
	for l < h {
		mid := l + (h-l+1)/2
		if mid*mid > x {
			h = mid - 1
		} else {
			l = mid
		}
	}
	return l
}