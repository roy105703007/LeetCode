func hammingDistance(x int, y int) int {
	z := x ^ y
	fmt.Println(z)
	cnt := 0
	for z != 0 {
		if z&1 == 1 {
			cnt++
		}
		z = z >> 1
	}
	return cnt
}