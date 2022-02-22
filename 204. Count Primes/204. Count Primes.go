func countPrimes(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 0
	}
	count := 0
	notPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		if notPrime[i] {
			continue
		}
		count++
		for k := i * i; k < n; k += i {
			notPrime[k] = true

		}
	}
	return count
}