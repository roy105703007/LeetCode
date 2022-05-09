func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	max := 0
	leftLow := false
	for i < j {
		leftLow = height[i] == min(height[i], height[j])
		if (j-i)*min(height[i], height[j]) > max {
			max = (j - i) * min(height[i], height[j])
		}
		if leftLow {
			i++
		} else {
			j--
		}
	}
	return max
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}