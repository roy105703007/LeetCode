func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 1 {
		if n == 0 {
			return true
		} else if n == 1 && flowerbed[0] == 0 {
			return true
		} else {
			return false
		}
	}
	if flowerbed[0] == 0 && flowerbed[1] == 0 {
		n -= 1
		flowerbed[0] = 1
	}
	for i := 1; i < len(flowerbed)-1; i++ {
		if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			n -= 1
		}
	}
	if flowerbed[len(flowerbed)-1] == 0 && flowerbed[len(flowerbed)-2] == 0 {
		n -= 1
	}
	return n <= 0
}